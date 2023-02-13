package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type Client struct {
	Logger    zerolog.Logger
	SCWClient *scw.Client
	Backend   backend.Backend
	OrgID     string
	Region    scw.Region
	Zone      scw.Zone

	Spec       Spec
	sourceSpec specs.Source
}

func (c *Client) ID() string {
	return c.sourceSpec.Name + ":" + string(c.Region) + ":" + string(c.Zone)
}

func (c *Client) WithRegion(r scw.Region) *Client {
	return &Client{
		Logger:     c.Logger.With().Str("region", string(r)).Logger(),
		SCWClient:  c.SCWClient,
		Backend:    c.Backend,
		OrgID:      c.OrgID,
		Region:     r,
		Zone:       c.Zone,
		Spec:       c.Spec,
		sourceSpec: c.sourceSpec,
	}
}

func (c *Client) WithZone(z scw.Zone) *Client {
	return &Client{
		Logger:     c.Logger.With().Str("zone", string(z)).Logger(),
		SCWClient:  c.SCWClient,
		Backend:    c.Backend,
		OrgID:      c.OrgID,
		Region:     c.Region,
		Zone:       z,
		Spec:       c.Spec,
		sourceSpec: c.sourceSpec,
	}
}

func New(_ context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec
	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
	err := pluginSpec.Validate()
	if err != nil {
		return nil, fmt.Errorf("failed to validate plugin spec: %w", err)
	}
	pluginSpec.SetDefaults()

	scwOpts := []scw.ClientOption{
		scw.WithEnv(), // existing env variables may overwrite active profile

		scw.WithHTTPClient(&http.Client{
			Timeout: time.Duration(pluginSpec.Timeout) * time.Second,
		}),
		scw.WithUserAgent("cq-plugin-scaleway/" + s.Version),
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
		Logger:     logger,
		Backend:    opts.Backend,
		OrgID:      orgID,
		Spec:       pluginSpec,
		SCWClient:  scwClient,
		sourceSpec: s,
	}, nil
}
