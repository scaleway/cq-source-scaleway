package mnq

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	mnq "github.com/scaleway/scaleway-sdk-go/api/mnq/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_mnq_namespaces",
		Resolver:  fetchNamespaces,
		Transform: transformers.TransformWithStruct(&mnq.Namespace{}, transformers.WithPrimaryKeys("ID")),
		Relations: []*schema.Table{
			namespaceCredentials(),
		},
	}
}

func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := mnq.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListNamespaces(&mnq.ListNamespacesRequest{
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
