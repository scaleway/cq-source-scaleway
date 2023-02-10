package lb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func lbPrivateNetworks() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lb_private_networks",
		Resolver:  fetchLBPrivateNetworks,
		Transform: transformers.TransformWithStruct(&lb.PrivateNetwork{}, transformers.WithPrimaryKeys("PrivateNetworkID")),
		Columns: schema.ColumnList{
			{
				Name:     "lb_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchLBPrivateNetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*lb.LB)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListLBPrivateNetworks(&lb.ListLBPrivateNetworksRequest{
			LBID:     p.ID,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.PrivateNetwork
		if len(response.PrivateNetwork) < int(limit) {
			break
		}

		page++
	}

	return nil
}
