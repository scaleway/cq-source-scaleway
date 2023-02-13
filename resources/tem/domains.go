package tem

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	tem "github.com/scaleway/scaleway-sdk-go/api/tem/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_tem_domains",
		Resolver:  fetchDomains,
		Transform: transformers.TransformWithStruct(&tem.Domain{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplex,
		Relations: []*schema.Table{
			domainEmails(),
		},
	}
}

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := tem.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListDomains(&tem.ListDomainsRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Domains
		if len(response.Domains) < int(limit) {
			break
		}

		page++
	}

	return nil
}
