package secrets

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/secret/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func secretVersions() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_secret_versions",
		Resolver:  fetchSecretVersions,
		Transform: transformers.TransformWithStruct(&secret.SecretVersion{}, transformers.WithPrimaryKeys("SecretID", "Revision")),
	}
}

func fetchSecretVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*secret.Secret)
	api := secret.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListSecretVersions(&secret.ListSecretVersionsRequest{
			SecretID: p.ID,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Versions
		if len(response.Versions) < int(limit) {
			break
		}

		page++
	}

	return nil
}
