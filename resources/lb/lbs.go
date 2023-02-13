package lb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func LBs() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lbs",
		Resolver:  fetchLBs,
		Transform: transformers.TransformWithStruct(&lb.LB{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplex,
		Relations: []*schema.Table{
			lbBackends(),
			lbBackendStats(),
			lbCertificates(),
			lbFrontends(),
			lbPrivateNetworks(),
		},
	}
}

func fetchLBs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListLBs(&lb.ListLBsRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.LBs
		if len(response.LBs) < int(limit) {
			break
		}

		page++
	}

	return nil
}
