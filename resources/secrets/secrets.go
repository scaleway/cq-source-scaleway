package secrets

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/secret/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_secrets",
		Resolver:  fetchSecrets,
		Transform: transformers.TransformWithStruct(&secret.Secret{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplex,
		Relations: []*schema.Table{
			secretVersions(),
		},
	}
}

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := secret.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListSecrets(&secret.ListSecretsRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Secrets
		if len(response.Secrets) < int(limit) {
			break
		}

		page++
	}

	return nil
}
