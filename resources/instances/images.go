package instances

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_images",
		Resolver:  fetchImages,
		Transform: transformers.TransformWithStruct(&instance.Image{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplex,
	}
}

func fetchImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := instance.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListImages(&instance.ListImagesRequest{
			Zone:         cl.Zone,
			Organization: &cl.OrgID,
			PerPage:      &limit,
			Page:         &page,
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
