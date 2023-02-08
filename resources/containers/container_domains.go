package containers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	container "github.com/scaleway/scaleway-sdk-go/api/container/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func containerDomains() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_container_domains",
		Resolver:  fetchContainerDomains,
		Transform: transformers.TransformWithStruct(&container.Domain{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchContainerDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*container.Container)
	api := container.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDomains(&container.ListDomainsRequest{
			ContainerID: p.ID,
			PageSize:    &limit,
			Page:        &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Domains
		if len(response.Domains) < int(limit) {
			break
		}

		page++
	}

	return nil
}
