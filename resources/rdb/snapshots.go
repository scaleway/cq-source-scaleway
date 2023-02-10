package rdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_snapshots",
		Resolver:  fetchSnapshots,
		Transform: transformers.TransformWithStruct(&rdb.Snapshot{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListSnapshots(&rdb.ListSnapshotsRequest{
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Snapshots
		if len(response.Snapshots) < int(limit) {
			break
		}

		page++
	}

	return nil
}
