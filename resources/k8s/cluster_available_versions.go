package k8s

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func clusterAvailableVersions() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_k8s_cluster_available_versions",
		Resolver:  fetchClusterAvailableVersions,
		Transform: transformers.TransformWithStruct(&k8s.Version{}, transformers.WithPrimaryKeys("Name")),
		Columns: schema.ColumnList{
			{
				Name:     "cluster_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}

func fetchClusterAvailableVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*k8s.Cluster)
	api := k8s.NewAPI(cl.SCWClient)

	response, err := api.ListClusterAvailableVersions(&k8s.ListClusterAvailableVersionsRequest{
		ClusterID: p.ID,
	}, scw.WithContext(ctx))
	if err != nil {
		return err
	}

	res <- response.Versions

	return nil
}
