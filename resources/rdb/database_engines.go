package rdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func DatabaseEngines() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_database_engines",
		Resolver:  fetchDatabaseEngines,
		Transform: transformers.TransformWithStruct(&rdb.DatabaseEngine{}, transformers.WithPrimaryKeys("Name", "Region")),
		Multiplex: client.RegionMultiplex,
	}
}

func fetchDatabaseEngines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDatabaseEngines(&rdb.ListDatabaseEnginesRequest{
			Region:   cl.Region,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Engines
		if len(response.Engines) < int(limit) {
			break
		}

		page++
	}

	return nil
}
