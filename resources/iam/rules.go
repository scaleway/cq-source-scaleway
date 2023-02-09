package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	iam "github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Rules() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_iam_rules",
		Resolver:  fetchRules,
		Transform: transformers.TransformWithStruct(&iam.Rule{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := iam.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListRules(&iam.ListRulesRequest{
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Rules
		if len(response.Rules) < int(limit) {
			break
		}

		page++
	}

	return nil
}
