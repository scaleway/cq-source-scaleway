package instances

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SecurityGroups() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_security_groups",
		Resolver:  fetchSecurityGroups,
		Transform: transformers.TransformWithStruct(&instance.SecurityGroup{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplex,
		Relations: []*schema.Table{
			securityGroupRules(),
		},
	}
}

func fetchSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := instance.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListSecurityGroups(&instance.ListSecurityGroupsRequest{
			Zone:         cl.Zone,
			Organization: &cl.OrgID,
			PerPage:      &limit,
			Page:         &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.SecurityGroups
		if len(response.SecurityGroups) < int(limit) {
			break
		}

		page++
	}

	return nil
}
