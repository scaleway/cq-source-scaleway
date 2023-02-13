package mnq

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	mnq "github.com/scaleway/scaleway-sdk-go/api/mnq/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func namespaceCredentials() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_mnq_namespace_credentials",
		Resolver:  fetchCredentials,
		Transform: transformers.TransformWithStruct(&mnq.Credential{}, transformers.WithPrimaryKeys("ID", "NamespaceID")),
	}
}

func fetchCredentials(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*mnq.Namespace)
	api := mnq.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListCredentials(&mnq.ListCredentialsRequest{
			Region:      p.Region,
			NamespaceID: &p.ID,
			PageSize:    &limit,
			Page:        &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Credentials
		if len(response.Credentials) < int(limit) {
			break
		}

		page++
	}

	return nil
}
