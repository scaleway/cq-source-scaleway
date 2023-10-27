package lb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func IPs() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lb_ips",
		Resolver:  fetchLBIPs,
		Transform: transformers.TransformWithStruct(&lb.IP{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplex,
	}
}

func fetchLBIPs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListIPs(&lb.ListIPsRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.IPs
		if len(response.IPs) < int(limit) {
			break
		}

		page++
	}

	return nil
}
