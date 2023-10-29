package ipfs

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/scaleway/cq-source-scaleway/client"
	ipfs "github.com/scaleway/scaleway-sdk-go/api/ipfs/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func volumePins() *schema.Table {
	return &schema.Table{
		Name:      "scaleway_ipfs_volume_pins",
		Resolver:  fetchVolumePins,
		Transform: transformers.TransformWithStruct(&ipfs.Pin{}, transformers.WithPrimaryKeys("PinID")),
		Columns: []schema.Column{
			{
				Name:       "volume_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchVolumePins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*ipfs.Volume)
	api := ipfs.NewAPI(cl.SCWClient)

	limit := uint32(100)
	page := int32(1)

	for {
		response, err := api.ListPins(&ipfs.ListPinsRequest{
			Region:   p.Region,
			VolumeID: p.ID,
			PageSize: &limit,
			Page:     &page,
		}, scw.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- response.Pins
		if len(response.Pins) < int(limit) {
			break
		}

		page++
	}

	return nil
}
