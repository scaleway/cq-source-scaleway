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
	val, ok := cl.SCWClient.GetDefaultZone()
	if !ok {
		return r.Set(c.Name, nil)
	}
	return r.Set(c.Name, val.String())
}

func ResolveClientRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*Client)
	val, ok := cl.SCWClient.GetDefaultRegion()
	if !ok {
		return r.Set(c.Name, nil)
	}
	return r.Set(c.Name, val.String())
}
