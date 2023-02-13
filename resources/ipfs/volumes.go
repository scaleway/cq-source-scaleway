package ipfs

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	ipfs "github.com/scaleway/scaleway-sdk-go/api/ipfs/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Volumes() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_ipfs_volumes",
		Resolver:  fetchVolumes,
		Transform: transformers.TransformWithStruct(&ipfs.Volume{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplex,
		Relations: []*schema.Table{
			volumePins(),
		},
	}
}

func fetchVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := ipfs.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListVolumes(&ipfs.ListVolumesRequest{
			Region:   cl.Region,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Volumes
		if len(response.Volumes) < int(limit) {
			break
		}

		page++
	}

	return nil
}
