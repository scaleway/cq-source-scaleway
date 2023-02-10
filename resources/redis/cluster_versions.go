package redis

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/redis/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Versions() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_redis_versions",
		Resolver:  fetchVersions,
		Transform: transformers.TransformWithStruct(&redis.ClusterVersion{}, transformers.WithPrimaryKeys("Version", "Zone")),
	}
}

func fetchVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := redis.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListClusterVersions(&redis.ListClusterVersionsRequest{
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Versions
		if len(response.Versions) < int(limit) {
			break
		}

		page++
	}

	return nil
}
