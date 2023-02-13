package functions

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func functionTriggers() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_function_triggers",
		Resolver:  fetchFunctionTriggers,
		Transform: transformers.TransformWithStruct(&function.Trigger{}, transformers.WithPrimaryKeys("ID")),
		Relations: []*schema.Table{
			functionTriggerInputs(),
		},
	}
}

func fetchFunctionTriggers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*function.Function)
	api := function.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListTriggers(&function.ListTriggersRequest{
			Region:     p.Region,
			FunctionID: p.ID,
			PageSize:   &limit,
			Page:       &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Triggers
		if len(response.Triggers) < int(limit) {
			break
		}

		page++
	}

	return nil
}
