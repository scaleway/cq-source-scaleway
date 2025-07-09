package rdb

import (
	"context"

	"github.com/apache/arrow/go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func instancePrivileges() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_instance_privileges",
		Resolver:  fetchInstancePrivileges,
		Transform: transformers.TransformWithStruct(&rdb.Privilege{}, transformers.WithPrimaryKeys("Permission", "DatabaseName", "UserName")),
		Columns: schema.ColumnList{
			{
				Name:     "instance_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchInstancePrivileges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*rdb.Instance)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListPrivileges(&rdb.ListPrivilegesRequest{
			Region:     p.Region,
			InstanceID: p.ID,
			PageSize:   &limit,
			Page:       &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Privileges
		if len(response.Privileges) < int(limit) {
			break
		}

		page++
	}

	return nil
}
