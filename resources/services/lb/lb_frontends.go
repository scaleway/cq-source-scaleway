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

func lbFrontends() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lb_frontends",
		Resolver:  fetchLBFrontends,
		Transform: transformers.TransformWithStruct(&lb.Frontend{}, transformers.WithPrimaryKeys("ID")),
		Columns: schema.ColumnList{
			{
				Name:       "lb_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			lbFrontendACLs(),
			lbFrontendRoutes(),
		},
	}
}

func fetchLBFrontends(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
		response, err := api.ListFrontends(&lb.ListFrontendsRequest{
			Region:   reg,
			LBID:     p.ID,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Frontends
		if len(response.Frontends) < int(limit) {
			break
		}

		page++
	}

	return nil
}
