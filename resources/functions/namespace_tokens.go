package functions

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func namespaceTokens() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_function_tokens",
		Resolver:  fetchTokens,
		Transform: transformers.TransformWithStruct(&function.Token{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchTokens(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*function.Namespace)
	api := function.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListTokens(&function.ListTokensRequest{
			Region:      p.Region,
			NamespaceID: &p.ID,
			PageSize:    &limit,
			Page:        &page,
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
