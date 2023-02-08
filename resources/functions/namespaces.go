package functions

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_function_namespaces",
		Resolver:  fetchNamespaces,
		Transform: transformers.TransformWithStruct(&function.Namespace{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := function.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListNamespaces(&function.ListNamespacesRequest{
			PageSize: &limit,
			Page:     &page,
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
