package functions

import (
	"context"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func functionTriggerInputs() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_function_trigger_inputs",
		Resolver:  fetchFunctionTriggerInputs,
		Transform: transformers.TransformWithStruct(&function.TriggerInput{}, transformers.WithPrimaryKeys("ID")),
		Columns: schema.ColumnList{
			{
				Name:       "trigger_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchFunctionTriggerInputs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*function.Trigger)
	pp := parent.Parent.Item.(*function.Function)
	api := function.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListTriggerInputs(&function.ListTriggerInputsRequest{
			Region:    pp.Region,
			TriggerID: p.ID,
			PageSize:  &limit,
			Page:      &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Inputs
		if len(response.Inputs) < int(limit) {
			break
		}

		page++
	}

	return nil
}
