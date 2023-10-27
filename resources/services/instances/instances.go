package instances

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_instances",
		Resolver:  fetchInstances,
		Transform: transformers.TransformWithStruct(&instance.Server{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplex,
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := instance.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListServers(&instance.ListServersRequest{
			Zone:         cl.Zone,
			Organization: &cl.OrgID,
			PerPage:      &limit,
			Page:         &page,
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
