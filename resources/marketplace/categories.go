package marketplace

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/marketplace/v2"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Categories() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_marketplace_categories",
		Resolver:  fetchCategories,
		Transform: transformers.TransformWithStruct(&marketplace.Category{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchCategories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := marketplace.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListCategories(&marketplace.ListCategoriesRequest{
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Categories
		if len(response.Categories) < int(limit) {
			break
		}

		page++
	}

	return nil
}
