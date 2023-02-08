package applesilicon

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	applesilicon "github.com/scaleway/scaleway-sdk-go/api/applesilicon/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func OS() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_applesilicon_os",
		Resolver:  fetchOS,
		Transform: transformers.TransformWithStruct(&applesilicon.OS{}, transformers.WithPrimaryKeys("ID")),
		Columns: schema.ColumnList{
			client.ZonePK,
		},
	}
}

func fetchOS(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := applesilicon.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListOS(&applesilicon.ListOSRequest{
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Os
		if len(response.Os) < int(limit) {
			break
		}

		page++
	}

	return nil
}
