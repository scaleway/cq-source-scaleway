package rdb

import (
	"context"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func instanceUsers() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_users",
		Resolver:  fetchUsers,
		Transform: transformers.TransformWithStruct(&rdb.User{}, transformers.WithPrimaryKeys("Name")),
		Columns: schema.ColumnList{
			{
				Name:     "instance_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*rdb.Instance)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListUsers(&rdb.ListUsersRequest{
			Region:     p.Region,
			InstanceID: p.ID,
			PageSize:   &limit,
			Page:       &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Users
		if len(response.Users) < int(limit) {
			break
		}

		page++
	}

	return nil
}
