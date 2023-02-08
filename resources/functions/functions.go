package functions

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Functions() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_functions",
		Resolver:  fetchFunctions,
		Transform: transformers.TransformWithStruct(&function.Function{}, transformers.WithPrimaryKeys("ID")),
		Relations: []*schema.Table{
			functionCrons(),
			functionDomains(),
			functionTriggers(),
		},
	}
}

func fetchFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := function.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListFunctions(&function.ListFunctionsRequest{
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Functions
		if len(response.Functions) < int(limit) {
			break
		}

		page++
	}

	return nil
}
