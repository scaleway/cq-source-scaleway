package rdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func instanceDatabases() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_databases",
		Resolver:  fetchDatabases,
		Transform: transformers.TransformWithStruct(&rdb.Database{}, transformers.WithPrimaryKeys("Name")),
		Columns: schema.ColumnList{
			{
				Name:     "instance_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*rdb.Instance)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDatabases(&rdb.ListDatabasesRequest{
			InstanceID: p.ID,
			PageSize:   &limit,
			Page:       &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Databases
		if len(response.Databases) < int(limit) {
			break
		}

		page++
	}

	return nil
}
