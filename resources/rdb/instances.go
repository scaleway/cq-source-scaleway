package rdb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_rdb_instances",
		Resolver:  fetchInstances,
		Transform: transformers.TransformWithStruct(&rdb.Instance{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.OrgRegionMultiplex,
		Relations: []*schema.Table{
			instanceACLRules(),
			instanceDatabases(),
			instancePrivileges(),
			instanceUsers(),
		},
	}
}

func fetchInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := rdb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListInstances(&rdb.ListInstancesRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Instances
		if len(response.Instances) < int(limit) {
			break
		}

		page++
	}

	return nil
}
