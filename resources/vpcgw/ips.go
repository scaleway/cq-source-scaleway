package vpcgw

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/vpcgw/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func IPs() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_vpcgw_ips",
		Resolver:  fetchIPs,
		Transform: transformers.TransformWithStruct(&vpcgw.IP{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchIPs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := vpcgw.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListIPs(&vpcgw.ListIPsRequest{
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.IPs
		if len(response.IPs) < int(limit) {
			break
		}

		page++
	}

	return nil
}
