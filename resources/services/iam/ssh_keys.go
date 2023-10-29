package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	iam "github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func SSHKeys() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_iam_ssh_keys",
		Resolver:  fetchSSHKeys,
		Transform: transformers.TransformWithStruct(&iam.SSHKey{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchSSHKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := iam.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListSSHKeys(&iam.ListSSHKeysRequest{
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.SSHKeys
		if len(response.SSHKeys) < int(limit) {
			break
		}

		page++
	}

	return nil
}
