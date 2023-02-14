package instances

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func securityGroupRules() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_security_group_rules",
		Resolver:  fetchSecurityGroupRules,
		Transform: transformers.TransformWithStruct(&instance.SecurityGroupRule{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			{
				Name:     "security_group_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchSecurityGroupRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*instance.SecurityGroup)
	api := instance.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListSecurityGroupRules(&instance.ListSecurityGroupRulesRequest{
			Zone:            p.Zone,
			SecurityGroupID: p.ID,
			PerPage:         &limit,
			Page:            &page,
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
