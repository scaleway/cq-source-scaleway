package redis

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/redis/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func NodeTypes() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_redis_node_types",
		Resolver:  fetchNodeTypes,
		Transform: transformers.TransformWithStruct(&redis.NodeType{}, transformers.WithPrimaryKeys("Name", "Zone")),
		Multiplex: client.ZoneMultiplexService("redis"),
	}
}

func fetchNodeTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := redis.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListNodeTypes(&redis.ListNodeTypesRequest{
			Zone:     cl.Zone,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.NodeTypes
		if len(response.NodeTypes) < int(limit) {
			break
		}

		page++
	}

	return nil
}
