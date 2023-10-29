package rdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func DatabaseBackups() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_database_backups",
		Resolver:  fetchDatabaseBackups,
		Transform: transformers.TransformWithStruct(&rdb.DatabaseBackup{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplex,
	}
}

func fetchDatabaseBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDatabaseBackups(&rdb.ListDatabaseBackupsRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.DatabaseBackups
		if len(response.DatabaseBackups) < int(limit) {
			break
		}

		page++
	}

	return nil
}
