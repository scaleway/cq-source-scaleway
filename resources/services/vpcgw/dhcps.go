package vpcgw

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/vpcgw/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func DHCPs() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_vpcgw_dhcps",
		Resolver:  fetchDHCPs,
		Transform: transformers.TransformWithStruct(&vpcgw.DHCP{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplexService("vpcgw"),
	}
}

func fetchDHCPs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := vpcgw.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDHCPs(&vpcgw.ListDHCPsRequest{
			Zone:           cl.Zone,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Dhcps
		if len(response.Dhcps) < int(limit) {
			break
		}

		page++
	}

	return nil
}
