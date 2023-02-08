package client

import (
	"context"
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

	Spec       Spec
	sourceSpec specs.Source
}

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

	cf, err := scw.LoadConfig()
	if err != nil {
		return nil, err
	}
	p, err := cf.GetActiveProfile()
	if err != nil {
		return nil, err
	}

	// Create a Scaleway client
	scwClient, err := scw.NewClient(
		// Get your credentials at https://console.scaleway.com/project/credentials
		scw.WithProfile(p), // active profile applies first
		scw.WithEnv(),      // existing env variables may overwrite active profile

		scw.WithHTTPClient(&http.Client{
			Timeout: time.Duration(pluginSpec.Timeout) * time.Second,
		}),
		scw.WithUserAgent("cq-plugin-scaleway/"+s.Version),
	)

	return &Client{
		Logger:     logger,
		Backend:    opts.Backend,
		Spec:       pluginSpec,
		SCWClient:  scwClient,
		sourceSpec: s,
	}, nil
}
