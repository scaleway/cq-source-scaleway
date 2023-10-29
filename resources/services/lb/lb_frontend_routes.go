package lb

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func lbFrontendRoutes() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lb_frontend_routes",
		Resolver:  fetchLBFrontendRoutes,
		Transform: transformers.TransformWithStruct(&lb.Route{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchLBFrontendRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*lb.Frontend)
	pp := parent.Parent.Item.(*lb.LB)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	reg, err := pp.Zone.Region()
	if err != nil {
		return fmt.Errorf("invalid region: %w", err)
	}

	for {
		response, err := api.ListRoutes(&lb.ListRoutesRequest{
			Region:     reg,
			FrontendID: &p.ID,
			PageSize:   &limit,
			Page:       &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Routes
		if len(response.Routes) < int(limit) {
			break
		}

		page++
	}

	return nil
}
