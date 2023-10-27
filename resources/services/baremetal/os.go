package baremetal

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func OS() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_baremetal_os",
		Resolver:  fetchOS,
		Transform: transformers.TransformWithStruct(&baremetal.OS{}, transformers.WithPrimaryKeys("ID")),
		Columns: schema.ColumnList{
			client.ZonePK,
		},
		Multiplex: client.ZoneMultiplexService("baremetal"),
	}
}

func fetchOS(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := baremetal.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListOS(&baremetal.ListOSRequest{
			Zone:     cl.Zone,
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
