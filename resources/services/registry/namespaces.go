package registry

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/registry/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_registry_namespaces",
		Resolver:  fetchNamespaces,
		Transform: transformers.TransformWithStruct(&registry.Namespace{}, transformers.WithPrimaryKeys("ID", "Region")),
		Multiplex: client.RegionMultiplex,
	}
}

func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := registry.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListNamespaces(&registry.ListNamespacesRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Namespaces
		if len(response.Namespaces) < int(limit) {
			break
		}

		page++
	}

	return nil
}
