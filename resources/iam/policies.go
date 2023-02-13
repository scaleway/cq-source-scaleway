package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	iam "github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Policies() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_iam_policies",
		Resolver:  fetchPolicies,
		Transform: transformers.TransformWithStruct(&iam.Policy{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.OrgMultiplex,
		Relations: []*schema.Table{
			policyRules(),
		},
	}
}

func fetchPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := iam.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListPolicies(&iam.ListPoliciesRequest{
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Policies
		if len(response.Policies) < int(limit) {
			break
		}

		page++
	}

	return nil
}
