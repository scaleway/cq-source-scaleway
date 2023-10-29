package vpc

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/vpc/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func PrivateNetworks() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_vpc_private_networks",
		Resolver:  fetchVPCs,
		Transform: transformers.TransformWithStruct(&vpc.PrivateNetwork{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplex,
	}
}

func fetchVPCs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := vpc.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListPrivateNetworks(&vpc.ListPrivateNetworksRequest{
			Zone:           cl.Zone,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.PrivateNetworks
		if len(response.PrivateNetworks) < int(limit) {
			break
		}

		page++
	}

	return nil
}
