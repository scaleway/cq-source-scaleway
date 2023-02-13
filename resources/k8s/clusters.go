package k8s

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_k8s_clusters",
		Resolver:  fetchClusters,
		Transform: transformers.TransformWithStruct(&k8s.Cluster{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.OrgRegionMultiplex,
		Relations: []*schema.Table{
			clusterAvailableVersions(),
			clusterNodes(),
			clusterPools(),
		},
	}
}

func fetchClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := k8s.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListClusters(&k8s.ListClustersRequest{
			Region:         cl.Region,
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
