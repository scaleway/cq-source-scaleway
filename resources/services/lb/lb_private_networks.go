package lb

import (
	"context"
	"fmt"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "lb_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
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

	reg, err := p.Zone.Region()
	if err != nil {
		return fmt.Errorf("invalid region: %w", err)
	}

	for {
		response, err := api.ListLBPrivateNetworks(&lb.ListLBPrivateNetworksRequest{
			Region:   reg,
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
