package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_baremetal_servers",
		Resolver:  fetchServers,
		Transform: transformers.TransformWithStruct(&baremetal.Server{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplexService("baremetal"),
	}
}

func fetchServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := baremetal.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListServers(&baremetal.ListServersRequest{
			Zone:           cl.Zone,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Servers
		if len(response.Servers) < int(limit) {
			break
		}

		page++
	}

	return nil
}
