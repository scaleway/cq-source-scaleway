package vpcgw

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/vpcgw/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func gatewayNetworks() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_vpcgw_gateway_networks",
		Resolver:  fetchGatewayNetworks,
		Transform: transformers.TransformWithStruct(&vpcgw.GatewayNetwork{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchGatewayNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*vpcgw.Gateway)
	api := vpcgw.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListGatewayNetworks(&vpcgw.ListGatewayNetworksRequest{
			GatewayID: &p.ID,
			PageSize:  &limit,
			Page:      &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.GatewayNetworks
		if len(response.GatewayNetworks) < int(limit) {
			break
		}

		page++
	}

	return nil
}
