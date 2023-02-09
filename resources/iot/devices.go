package iot

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/iot/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Devices() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_iot_devices",
		Resolver:  fetchDevices,
		Transform: transformers.TransformWithStruct(&iot.Device{}, transformers.WithPrimaryKeys("ID")),
		Columns: schema.ColumnList{
			client.RegionPK,
		},
	}
}

func fetchDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := iot.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDevices(&iot.ListDevicesRequest{
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Devices
		if len(response.Devices) < int(limit) {
			break
		}

		page++
	}

	return nil
}
