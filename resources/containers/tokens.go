package containers

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	container "github.com/scaleway/scaleway-sdk-go/api/container/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Tokens() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_container_tokens",
		Resolver:  fetchTokens,
		Transform: transformers.TransformWithStruct(&container.Token{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchTokens(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := container.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListTokens(&container.ListTokensRequest{
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Tokens
		if len(response.Tokens) < int(limit) {
			break
		}

		page++
	}

	return nil
}
