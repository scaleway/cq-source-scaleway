package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

var TestOrgID = "DECAF000-CAFE-0000-0000-000000000000"

func TestHelper(t *testing.T, table *schema.Table, createServices func(*mux.Router) error) {
	version := "vDev"

	t.Helper()
	table.IgnoreInTests = false

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Router received request to %s", r.URL.String())
		http.Error(w, "not found", http.StatusNotFound)
	})

	h := httptest.NewServer(router)
	defer h.Close()

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, opts source.Options) (schema.ClientMeta, error) {
		scwClient, err := scw.NewClient(
			scw.WithoutAuth(),
			scw.WithInsecure(),
			scw.WithAPIURL(h.URL),
			scw.WithDefaultOrganizationID(TestOrgID),
		)
		if err != nil {
			return nil, err
		}
		s := Spec{
			OrgIDs: []string{TestOrgID},
		}
		s.SetDefaults()
		if err := s.Validate(); err != nil {
			return nil, err
		}

		if err := createServices(router); err != nil {
			return nil, err
		}

		return &Client{
			Logger:     logger,
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
	p.SetLogger(logger)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
