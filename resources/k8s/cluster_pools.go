package k8s

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func clusterPools() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_k8s_pools",
		Resolver:  fetchClusterPools,
		Transform: transformers.TransformWithStruct(&k8s.Pool{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchClusterPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*k8s.Cluster)
	api := k8s.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListPools(&k8s.ListPoolsRequest{
			ClusterID: p.ID,
			PageSize:  &limit,
			Page:      &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Pools
		if len(response.Pools) < int(limit) {
			break
		}

		page++
	}

	return nil
}
