package vpcgw

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/vpcgw/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func DHCPEntries() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_vpcgw_dhcp_entries",
		Resolver:  fetchDHCPEntries,
		Transform: transformers.TransformWithStruct(&vpcgw.DHCPEntry{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplexService("vpcgw"),
	}
}

func fetchDHCPEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := vpcgw.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDHCPEntries(&vpcgw.ListDHCPEntriesRequest{
			Zone:     cl.Zone,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.DHCPEntries
		if len(response.DHCPEntries) < int(limit) {
			break
		}

		page++
	}

	return nil
}
