package registry

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/registry/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type ImageTag struct {
	ImageID string
	Tags    []*registry.Tag
}

func imageTags() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_registry_image_tags",
		Resolver:  fetchImageTags,
		Transform: transformers.TransformWithStruct(&ImageTag{}, transformers.WithPrimaryKeys("ImageID")),
	}
}

func fetchImageTags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*registry.Image)
	api := registry.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	obj := &ImageTag{
		ImageID: p.ID,
	}

	for {
		response, err := api.ListTags(&registry.ListTagsRequest{
			ImageID:  p.ID,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		obj.Tags = append(obj.Tags, response.Tags...)

		if len(response.Tags) < int(limit) {
			break
		}

		page++
	}

	if len(obj.Tags) > 0 {
		res <- obj
	}

	return nil
}
