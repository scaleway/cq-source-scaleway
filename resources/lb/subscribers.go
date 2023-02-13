package lb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Subscribers() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lb_subscribers",
		Resolver:  fetchSubscribers,
		Transform: transformers.TransformWithStruct(&lb.Subscriber{}, transformers.WithPrimaryKeys("ID")),
		Multiplex: client.RegionMultiplex,
	}
}

func fetchSubscribers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListSubscriber(&lb.ListSubscriberRequest{
			Region:         cl.Region,
			OrganizationID: &cl.OrgID,
			PageSize:       &limit,
			Page:           &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Subscribers
		if len(response.Subscribers) < int(limit) {
			break
		}

		page++
	}

	return nil
}
