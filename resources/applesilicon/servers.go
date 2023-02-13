package applesilicon

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	applesilicon "github.com/scaleway/scaleway-sdk-go/api/applesilicon/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_applesilicon_servers",
		Resolver:  fetchServers,
		Transform: transformers.TransformWithStruct(&applesilicon.Server{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.OrgZoneMultiplex,
	}
}

func fetchServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := applesilicon.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListServers(&applesilicon.ListServersRequest{
			Zone:           cl.Zone,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Servers
		if len(response.Servers) < int(limit) {
			break
		}

		page++
	}

	return nil
}
