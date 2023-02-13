package vpcgw

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/vpcgw/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Gateways() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_vpcgw_gateways",
		Resolver:  fetchGateways,
		Transform: transformers.TransformWithStruct(&vpcgw.Gateway{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.OrgZoneMultiplex,
		Relations: []*schema.Table{
			gatewayNetworks(),
		},
	}
}

func fetchGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := vpcgw.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListGateways(&vpcgw.ListGatewaysRequest{
			Zone:           cl.Zone,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Gateways
		if len(response.Gateways) < int(limit) {
			break
		}

		page++
	}

	return nil
}
