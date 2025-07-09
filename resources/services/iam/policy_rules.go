package iam

import (
	"context"

	"github.com/apache/arrow/go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	iam "github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func policyRules() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_iam_policy_rules",
		Resolver:  fetchPolicyRules,
		Transform: transformers.TransformWithStruct(&iam.Rule{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			{
				Name:       "policy_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchPolicyRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*iam.Policy)
	api := iam.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListRules(&iam.ListRulesRequest{
			PolicyID: &p.ID,
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
