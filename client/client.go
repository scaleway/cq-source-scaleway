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

	Spec       Spec
	sourceSpec specs.Source
}

const (
	defaultRegion = "fr-par"
	defaultZone   = "fr-par-1"
)

func (c *Client) ID() string {
	return c.sourceSpec.Name
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

	if _, ok := scwClient.GetDefaultOrganizationID(); !ok {
		return nil, fmt.Errorf("SCW_DEFAULT_ORGANIZATION_ID or default_organization_id not set, get yours from https://console.scaleway.com/organization/settings")
	}

	reinit := false
	if _, ok := scwClient.GetDefaultRegion(); !ok {
		logger.Log().Msg("SCW_DEFAULT_REGION or default_region not set, assuming " + defaultRegion)
		scwOpts = append(scwOpts, scw.WithDefaultRegion(defaultRegion))
		reinit = true
	}
	if _, ok := scwClient.GetDefaultZone(); !ok {
		logger.Log().Msg("SCW_DEFAULT_ZONE or default_zone not set, assuming " + defaultZone)
		scwOpts = append(scwOpts, scw.WithDefaultZone(defaultZone))
		reinit = true
	}

	if reinit {
		scwClient, err = scw.NewClient(scwOpts...)
		if err != nil {
			return nil, err
		}
	}

	orgID, _ := scwClient.GetDefaultOrganizationID()

	return &Client{
		Logger:     logger,
		Backend:    opts.Backend,
		OrgID:      orgID,
		Spec:       pluginSpec,
		SCWClient:  scwClient,
		sourceSpec: s,
	}, nil
}
