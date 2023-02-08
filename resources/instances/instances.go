package instances

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_instances",
		Resolver:  fetchInstances,
		Transform: transformers.TransformWithStruct(&instance.Server{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := instance.NewAPI(cl.SCWClient)

	for _, st := range []instance.ServerState{
		instance.ServerStateRunning,
		instance.ServerStateStarting,
		instance.ServerStateStopping,
		instance.ServerStateStopped,
		instance.ServerStateStoppedInPlace,
		instance.ServerStateLocked,
	} {
		limit := uint32(100)
		page := int32(1)

		for {
			response, err := api.ListServers(&instance.ListServersRequest{
				PerPage: &limit,
				Page:    &page,
				State:   &st,
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
	}

	return nil
}
