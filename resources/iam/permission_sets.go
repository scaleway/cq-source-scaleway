package iam

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	iam "github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func PermissionSets() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_iam_permission_sets",
		Resolver:  fetchPermissionSets,
		Transform: transformers.TransformWithStruct(&iam.PermissionSet{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.OrgMultiplex,
	}
}

func fetchPermissionSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := iam.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListPermissionSets(&iam.ListPermissionSetsRequest{
			OrganizationID: cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.PermissionSets
		if len(response.PermissionSets) < int(limit) {
			break
		}

		page++
	}

	return nil
}
