package instances

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func DefaultSecurityGroupRules() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_default_security_group_rules",
		Resolver:  fetchDefaultSecurityGroupRules,
		Transform: transformers.TransformWithStruct(&instance.SecurityGroupRule{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplex,
	}
}

func fetchDefaultSecurityGroupRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := instance.NewAPI(cl.SCWClient)

	response, err := api.ListDefaultSecurityGroupRules(&instance.ListDefaultSecurityGroupRulesRequest{
		Zone: cl.Zone,
	}, scw.WithContext(ctx))
	if err != nil {
		return err
	}

	res <- response.Rules

	return nil
}
