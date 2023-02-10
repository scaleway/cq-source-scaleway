package lb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "lb_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			lbFrontendACLs(),
		},
	}
}

func fetchLBFrontends(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*lb.LB)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListFrontends(&lb.ListFrontendsRequest{
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
