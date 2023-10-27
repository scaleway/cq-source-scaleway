package containers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	container "github.com/scaleway/scaleway-sdk-go/api/container/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func containerCrons() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_container_crons",
		Resolver:  fetchContainerCrons,
		Transform: transformers.TransformWithStruct(&container.Cron{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchContainerCrons(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*container.Container)
	api := container.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListCrons(&container.ListCronsRequest{
			Region:      p.Region,
			ContainerID: p.ID,
			PageSize:    &limit,
			Page:        &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Crons
		if len(response.Crons) < int(limit) {
			break
		}

		page++
	}

	return nil
}
