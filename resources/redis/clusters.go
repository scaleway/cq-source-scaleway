package redis

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/redis/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_redis_clusters",
		Resolver:  fetchClusters,
		Transform: transformers.TransformWithStruct(&redis.Cluster{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplexService("redis"),
	}
}

func fetchClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := redis.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListClusters(&redis.ListClustersRequest{
			Zone:           cl.Zone,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Clusters
		if len(response.Clusters) < int(limit) {
			break
		}

		page++
	}

	return nil
}
