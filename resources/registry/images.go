package registry

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/registry/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_registry_images",
		Resolver:  fetchImages,
		Transform: transformers.TransformWithStruct(&registry.Image{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplex,
		Relations: []*schema.Table{
			imageTags(),
		},
	}
}

func fetchImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := registry.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListImages(&registry.ListImagesRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Images
		if len(response.Images) < int(limit) {
			break
		}

		page++
	}

	return nil
}
