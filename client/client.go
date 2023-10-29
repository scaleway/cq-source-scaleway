package client

import (
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type Client struct {
	Logger    zerolog.Logger
	Spec      Spec
	SCWClient *scw.Client
	Backend   state.Client

	OrgID  string
	Region scw.Region
	Zone   scw.Zone
}

func New(logger zerolog.Logger, spec Spec, orgID string, services *scw.Client, bk state.Client) *Client {
	return &Client{
		Logger:    logger,
		Spec:      spec,
		OrgID:     orgID,
		SCWClient: services,
		Backend:   bk,
	}
}

func (c *Client) ID() string {
	return "scaleway:" + string(c.Region) + ":" + string(c.Zone)
}

func (c *Client) WithRegion(r scw.Region) *Client {
	return &Client{
		Logger:    c.Logger.With().Str("region", string(r)).Logger(),
		Spec:      c.Spec,
		SCWClient: c.SCWClient,
		Backend:   c.Backend,

		OrgID:  c.OrgID,
		Region: r,
		Zone:   c.Zone,
	}
}

func (c *Client) WithZone(z scw.Zone) *Client {
	return &Client{
		Logger:    c.Logger.With().Str("zone", string(z)).Logger(),
		Spec:      c.Spec,
		SCWClient: c.SCWClient,
		Backend:   c.Backend,

		OrgID:  c.OrgID,
		Region: c.Region,
		Zone:   z,
	}
}
