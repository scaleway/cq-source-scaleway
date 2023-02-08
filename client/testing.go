package client

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func TestHelper(t *testing.T, table *schema.Table, ts *httptest.Server) {
	version := "vDev"
	table.IgnoreInTests = false
	t.Helper()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, opts source.Options) (schema.ClientMeta, error) {
		scwClient, err := scw.NewClient(
			scw.WithoutAuth(),
			scw.WithInsecure(),
			scw.WithHTTPClient(ts.Client()),
		)
		if err != nil {
			return nil, err
		}
		s := Spec{}
		s.SetDefaults()
		if err := s.Validate(); err != nil {
			return nil, err
		}
		return &Client{
			Logger:     l,
			SCWClient:  scwClient,
			Backend:    opts.Backend,
			Spec:       s,
			sourceSpec: spec,
		}, nil
	}
	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
