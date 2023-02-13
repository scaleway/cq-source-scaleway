package rdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func instanceACLRules() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_instance_acl_rules",
		Resolver:  fetchInstanceACLRules,
		Transform: transformers.TransformWithStruct(&rdb.ACLRule{}, transformers.WithPrimaryKeys("IP", "Port", "Protocol", "Direction", "Action")),
		Columns: schema.ColumnList{
			{
				Name:     "instance_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchInstanceACLRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*rdb.Instance)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListInstanceACLRules(&rdb.ListInstanceACLRulesRequest{
			Region:     p.Region,
			InstanceID: p.ID,
			PageSize:   &limit,
			Page:       &page,
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
