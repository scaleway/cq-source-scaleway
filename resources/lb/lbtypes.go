package lb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func LBTypes() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lb_types",
		Resolver:  fetchLBTypes,
		Transform: transformers.TransformWithStruct(&lb.LBType{}, transformers.WithPrimaryKeys("Name", "Region")),
	}
}

func fetchLBTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListLBTypes(&lb.ListLBTypesRequest{
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.LBTypes
		if len(response.LBTypes) < int(limit) {
			break
		}

		page++
	}

	return nil
}
