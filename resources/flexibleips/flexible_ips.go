package flexibleips

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	flexibleip "github.com/scaleway/scaleway-sdk-go/api/flexibleip/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func FlexibleIPs() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_flexible_ips",
		Resolver:  fetchIPs,
		Transform: transformers.TransformWithStruct(&flexibleip.FlexibleIP{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.ZoneMultiplexService("flexibleip"),
	}
}

func fetchIPs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := flexibleip.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListFlexibleIPs(&flexibleip.ListFlexibleIPsRequest{
			Zone:           cl.Zone,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.FlexibleIPs
		if len(response.FlexibleIPs) < int(limit) {
			break
		}

		page++
	}

	return nil
}
