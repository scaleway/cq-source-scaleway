package applesilicon

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	applesilicon "github.com/scaleway/scaleway-sdk-go/api/applesilicon/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func ServerTypes() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_applesilicon_server_types",
		Resolver:  fetchServerTypes,
		Transform: transformers.TransformWithStruct(&applesilicon.ServerType{}, transformers.WithPrimaryKeys("Name")),
		Columns: schema.ColumnList{
			client.ZonePK,
		},
		Multiplex: client.ZoneMultiplex,
	}
}

func fetchServerTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := applesilicon.NewAPI(cl.SCWClient)

	response, err := api.ListServerTypes(&applesilicon.ListServerTypesRequest{
		Zone: cl.Zone,
	}, scw.WithContext(ctx))
	if err != nil {
		return err
	}

	res <- response.ServerTypes

	return nil
}
