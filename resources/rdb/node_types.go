package rdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func NodeTypes() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_node_types",
		Resolver:  fetchNodeTypes,
		Transform: transformers.TransformWithStruct(&rdb.NodeType{}, transformers.WithPrimaryKeys("Name", "Region")),
	}
}

func fetchNodeTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListNodeTypes(&rdb.ListNodeTypesRequest{
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
