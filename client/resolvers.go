package client

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	ZonePK = schema.Column{
		Name:       "zone",
		Type:       arrow.BinaryTypes.String,
		Resolver:   ResolveClientZone,
		PrimaryKey: true,
	}

	RegionPK = schema.Column{
		Name:       "region",
		Type:       arrow.BinaryTypes.String,
		Resolver:   ResolveClientRegion,
		PrimaryKey: true,
	}
)

func ResolveClientZone(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*Client)
	return r.Set(c.Name, cl.Zone)
}

func ResolveClientRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*Client)
	return r.Set(c.Name, cl.Region)
}
