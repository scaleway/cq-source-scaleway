package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	iam "github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func APIKeys() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_iam_api_keys",
		Resolver:  fetchAPIKeys,
		Transform: transformers.TransformWithStruct(&iam.APIKey{}, transformers.WithPrimaryKeys("AccessKey")),
		Multiplex: client.OrgMultiplex,
	}
}

func fetchAPIKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := iam.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListAPIKeys(&iam.ListAPIKeysRequest{
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.APIKeys
		if len(response.APIKeys) < int(limit) {
			break
		}

		page++
	}

	return nil
}
