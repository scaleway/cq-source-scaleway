package account

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/account/v2"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_projects",
		Resolver:  fetchProjects,
		Transform: transformers.TransformWithStruct(&account.Project{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := account.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListProjects(&account.ListProjectsRequest{
			OrganizationID: cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Projects
		if len(response.Projects) < int(limit) {
			break
		}

		page++
	}

	return nil
}
