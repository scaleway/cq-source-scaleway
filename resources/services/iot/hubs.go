package iot

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/iot/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Hubs() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_iot_hubs",
		Resolver:  fetchHubs,
		Transform: transformers.TransformWithStruct(&iot.Hub{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplexService("iot"),
		Relations: []*schema.Table{
			devices(),
			networks(),
			routes(),
		},
	}
}

func fetchHubs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := iot.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListHubs(&iot.ListHubsRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Hubs
		if len(response.Hubs) < int(limit) {
			break
		}

		page++
	}

	return nil
}
