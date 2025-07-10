package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/docs"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	internalPlugin "github.com/scaleway/cq-source-scaleway/plugin"
	"github.com/scaleway/scaleway-sdk-go/scw"

	"github.com/rs/zerolog"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/cq-source-scaleway/resources/services/account"
	"github.com/scaleway/cq-source-scaleway/resources/services/applesilicon"
	"github.com/scaleway/cq-source-scaleway/resources/services/baremetal"
	"github.com/scaleway/cq-source-scaleway/resources/services/containers"
	"github.com/scaleway/cq-source-scaleway/resources/services/flexibleips"
	"github.com/scaleway/cq-source-scaleway/resources/services/functions"
	"github.com/scaleway/cq-source-scaleway/resources/services/iam"
	"github.com/scaleway/cq-source-scaleway/resources/services/instances"
	"github.com/scaleway/cq-source-scaleway/resources/services/iot"
	"github.com/scaleway/cq-source-scaleway/resources/services/ipfs"
	"github.com/scaleway/cq-source-scaleway/resources/services/k8s"
	"github.com/scaleway/cq-source-scaleway/resources/services/lb"
	"github.com/scaleway/cq-source-scaleway/resources/services/marketplace"
	"github.com/scaleway/cq-source-scaleway/resources/services/mnq"
	"github.com/scaleway/cq-source-scaleway/resources/services/rdb"
	"github.com/scaleway/cq-source-scaleway/resources/services/redis"
	"github.com/scaleway/cq-source-scaleway/resources/services/registry"
	"github.com/scaleway/cq-source-scaleway/resources/services/secrets"
	"github.com/scaleway/cq-source-scaleway/resources/services/tem"
	"github.com/scaleway/cq-source-scaleway/resources/services/test"
	"github.com/scaleway/cq-source-scaleway/resources/services/vpc"
	"github.com/scaleway/cq-source-scaleway/resources/services/vpcgw"
)

type Client struct {
	logger    zerolog.Logger
	config    client.Spec
	tables    schema.Tables
	options   plugin.NewClientOptions
	scheduler *scheduler.Scheduler
	services  *scw.Client
	orgID     string

	plugin.UnimplementedDestination
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	stateClient, err := state.NewConnectedClient(ctx, options.BackendOptions)
	if err != nil {
		return err
	}
	defer stateClient.Close() //nolint:errcheck

	schedulerClient := client.New(c.logger, c.config, c.orgID, c.services, stateClient)
	err = c.scheduler.Sync(ctx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
	if err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return stateClient.Flush(ctx)
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (c *Client) Close(_ context.Context) error {
	return nil
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	if opts.NoConnection {
		return &Client{
			logger:  logger.With().Str("module", "scaleway").Logger(),
			options: opts,
			tables:  getTables(),
		}, nil
	}

	config := client.Spec{}
	if err := json.Unmarshal(specBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate spec: %w", err)
	}

	scwOpts := []scw.ClientOption{
		scw.WithEnv(), // existing env variables may overwrite active profile

		scw.WithHTTPClient(&http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		}),
		scw.WithUserAgent("cq-plugin-scaleway/" + internalPlugin.Version),
	}

	cf, err := scw.LoadConfig()
	if err != nil {
		var configNotFoundError *scw.ConfigFileNotFoundError
		if !errors.As(err, &configNotFoundError) {
			return nil, err
		}
	}
	if cf != nil {
		p, err := cf.GetActiveProfile()
		if err != nil {
			return nil, err
		}
		scwOpts = append([]scw.ClientOption{
			scw.WithProfile(p), // active profile applies first
		}, scwOpts...)
	}

	// Create a Scaleway client
	scwClient, err := scw.NewClient(scwOpts...)
	if err != nil {
		return nil, err
	}

	orgID, ok := scwClient.GetDefaultOrganizationID()
	if !ok {
		return nil, fmt.Errorf("SCW_DEFAULT_ORGANIZATION_ID or default_organization_id not set, get yours from https://console.scaleway.com/organization/settings")
	}

	return &Client{
		logger:  logger.With().Str("module", "scaleway").Logger(),
		options: opts,
		config:  config,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithConcurrency(config.Concurrency),
		),
		orgID:    orgID,
		services: scwClient,
		tables:   getTables(),
	}, nil
}

func getTables() schema.Tables {
	tables := []*schema.Table{
		account.Projects(),
		applesilicon.OS(),
		applesilicon.ServerTypes(),
		applesilicon.Servers(),
		baremetal.OS(),
		baremetal.Offers(),
		baremetal.Servers(),
		containers.Containers(),
		containers.Namespaces(),
		flexibleips.FlexibleIPs(),
		functions.Functions(),
		functions.Namespaces(),
		functions.Runtimes(),
		iam.APIKeys(),
		iam.Apps(),
		iam.Groups(),
		iam.PermissionSets(),
		iam.Policies(),
		iam.SSHKeys(),
		iam.Users(),
		instances.DefaultSecurityGroupRules(),
		instances.Images(),
		instances.Instances(),
		instances.SecurityGroups(),
		iot.Hubs(),
		ipfs.Volumes(),
		k8s.Clusters(),
		k8s.Versions(),
		lb.LBs(),
		lb.LBTypes(),
		lb.IPs(),
		lb.Subscribers(),
		marketplace.Categories(),
		marketplace.Images(),
		mnq.Namespaces(),
		rdb.DatabaseBackups(),
		rdb.DatabaseEngines(),
		rdb.Instances(),
		rdb.Snapshots(),
		rdb.NodeTypes(),
		redis.Clusters(),
		redis.Versions(),
		redis.NodeTypes(),
		registry.Images(),
		registry.Namespaces(),
		secrets.Secrets(),
		tem.Domains(),
		test.Humans(),
		vpc.PrivateNetworks(),
		vpcgw.DHCPs(),
		vpcgw.DHCPEntries(),
		vpcgw.Gateways(),
		vpcgw.IPs(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		panic(err)
	}
	if err := transformers.Apply(tables, func(t *schema.Table) error {
		t.Title = docs.DefaultTitleTransformer(t)
		return nil
	}); err != nil {
		panic(err)
	}
	for _, t := range tables {
		schema.AddCqIDs(t)
	}
	return tables
}
