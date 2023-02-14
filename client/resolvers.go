package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	ZonePK = schema.Column{
		Name:     "zone",
		Type:     schema.TypeString,
		Resolver: ResolveClientZone,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: true,
		},
	}

	RegionPK = schema.Column{
		Name:     "region",
		Type:     schema.TypeString,
		Resolver: ResolveClientRegion,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: true,
		},
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
