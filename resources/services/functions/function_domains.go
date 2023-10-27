package functions

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func functionDomains() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_function_domains",
		Resolver:  fetchFunctionDomains,
		Transform: transformers.TransformWithStruct(&function.Domain{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchFunctionDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*function.Function)
	api := function.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDomains(&function.ListDomainsRequest{
			Region:     p.Region,
			FunctionID: p.ID,
			PageSize:   &limit,
			Page:       &page,
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
