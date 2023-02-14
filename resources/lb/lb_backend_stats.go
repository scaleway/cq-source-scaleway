package lb

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func lbBackendStats() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_lb_backend_stats",
		Resolver:  fetchLBBackendStats,
		Transform: transformers.TransformWithStruct(&lb.BackendServerStats{}, transformers.WithPrimaryKeys("InstanceID", "BackendID", "IP")),
		Columns: schema.ColumnList{
			{
				Name:     "lb_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchLBBackendStats(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*lb.LB)
	api := lb.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListBackendStats(&lb.ListBackendStatsRequest{
			LBID:     p.ID,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.BackendServersStats
		if len(response.BackendServersStats) < int(limit) {
			break
		}

		page++
	}

	return nil
}
