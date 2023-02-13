package k8s

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Versions() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_k8s_versions",
		Resolver:  fetchVersions,
		Transform: transformers.TransformWithStruct(&k8s.Version{}, transformers.WithPrimaryKeys("Name", "Region")),
		Multiplex: client.RegionMultiplex,
	}
}

func fetchVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := k8s.NewAPI(cl.SCWClient)

	response, err := api.ListVersions(&k8s.ListVersionsRequest{
		Region: cl.Region,
	}, scw.WithContext(ctx))
	if err != nil {
		return err
	}

	res <- response.Versions

	return nil
}
