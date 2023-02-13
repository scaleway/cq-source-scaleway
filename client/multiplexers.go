package client

import "github.com/cloudquery/plugin-sdk/schema"

func OrgMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	cl := meta.(*Client)
	for _, o := range cl.Spec.OrgIDs {
		l = append(l, cl.WithOrg(o))
	}
	return l
}

func ZoneMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	cl := meta.(*Client)
	for _, z := range cl.Spec.Zones {
		l = append(l, cl.WithZone(z))
	}
	return l
}

func RegionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	cl := meta.(*Client)
	for _, r := range cl.Spec.Regions {
		l = append(l, cl.WithRegion(r))
	}
	return l
}

func OrgZoneMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	cl := meta.(*Client)
	for _, o := range cl.Spec.OrgIDs {
		for _, z := range cl.Spec.Zones {
			l = append(l, cl.WithOrg(o).WithZone(z))
		}
	}
	return l
}

func OrgRegionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	cl := meta.(*Client)
	for _, o := range cl.Spec.OrgIDs {
		for _, r := range cl.Spec.Regions {
			l = append(l, cl.WithOrg(o).WithRegion(r))
		}
	}
	return l
}
