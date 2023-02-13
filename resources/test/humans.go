package test

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/test/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Humans() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_test_humans",
		Resolver:  fetchHumans,
		Transform: transformers.TransformWithStruct(&test.Human{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.OrgMultiplex,
	}
}

func fetchHumans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := test.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListHumans(&test.ListHumansRequest{
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Humans
		if len(response.Humans) < int(limit) {
			break
		}

		page++
	}

	return nil
}
