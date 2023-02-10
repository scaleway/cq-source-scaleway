package lb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func lbFrontendACLs() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lb_frontend_acls",
		Resolver:  fetchLBFrontendACLs,
		Transform: transformers.TransformWithStruct(&lb.ACL{}, transformers.WithPrimaryKeys("ID")),
		Columns: schema.ColumnList{
			{
				Name:     "frontend_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchLBFrontendACLs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*lb.Frontend)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListACLs(&lb.ListACLsRequest{
			FrontendID: p.ID,
			Page:       &page,
			PageSize:   &limit,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.ACLs
		if len(response.ACLs) < int(limit) {
			break
		}

		page++
	}

	return nil
}
