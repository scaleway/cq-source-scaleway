package functions

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Runtimes() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_function_runtimes",
		Resolver:  fetchRuntimes,
		Transform: transformers.TransformWithStruct(&function.Runtime{}, transformers.WithPrimaryKeys("Name")),
		Columns: schema.ColumnList{
			client.RegionPK,
		},
		Multiplex: client.RegionMultiplex,
	}
}

func fetchRuntimes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := function.NewAPI(cl.SCWClient)

	response, err := api.ListFunctionRuntimes(&function.ListFunctionRuntimesRequest{
		Region: cl.Region,
	}, scw.WithContext(ctx))
	if err != nil {
		return err
	}

	res <- response.Runtimes

	return nil
}
