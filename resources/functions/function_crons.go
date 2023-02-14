package functions

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func functionCrons() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_function_crons",
		Resolver:  fetchFunctionCrons,
		Transform: transformers.TransformWithStruct(&function.Cron{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchFunctionCrons(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*function.Function)
	api := function.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListCrons(&function.ListCronsRequest{
			Region:     p.Region,
			FunctionID: p.ID,
			PageSize:   &limit,
			Page:       &page,
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
