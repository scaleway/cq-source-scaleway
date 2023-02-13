package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func OrgMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	var l = make([]schema.ClientMeta, 0)
	cl := meta.(*Client)
	for _, o := range cl.Spec.OrgIDs {
		l = append(l, cl.WithOrg(o))
	}
	return l
}

func ZoneMultiplexSelective(list ...scw.Zone) schema.Multiplexer {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		cl := meta.(*Client)
		for _, z := range list {
			l = append(l, cl.WithZone(z))
		}
		return l
	}
}

func ZoneMultiplexExcept(except ...scw.Zone) schema.Multiplexer {
	skipMap := make(map[scw.Zone]struct{}, len(except))
	for i := range except {
		skipMap[except[i]] = struct{}{}
	}

	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		cl := meta.(*Client)
		for _, z := range cl.Spec.Zones {
			if _, ok := skipMap[z]; !ok {
				l = append(l, cl.WithZone(z))
			}
		}
		return l
	}
}

func ZoneMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	return ZoneMultiplexExcept()(meta)
}

func RegionMultiplexSelective(list ...scw.Region) schema.Multiplexer {
	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		cl := meta.(*Client)
		for _, r := range list {
			l = append(l, cl.WithRegion(r))
		}
		return l
	}
}

func RegionMultiplexExcept(except ...scw.Region) schema.Multiplexer {
	skipMap := make(map[scw.Region]struct{}, len(except))
	for i := range except {
		skipMap[except[i]] = struct{}{}
	}

	return func(meta schema.ClientMeta) []schema.ClientMeta {
		var l = make([]schema.ClientMeta, 0)
		cl := meta.(*Client)
		for _, r := range cl.Spec.Regions {
			if _, ok := skipMap[r]; !ok {
				l = append(l, cl.WithRegion(r))
			}
		}
		return l
	}
}

func RegionMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	return RegionMultiplexExcept()(meta)
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

func ZoneMultiplexService(service string) schema.Multiplexer {
	switch service {
	case "applesilicon":
		return ZoneMultiplexSelective(scw.ZoneFrPar3)
	case "baremetal":
		return ZoneMultiplexSelective(scw.ZoneFrPar1, scw.ZoneFrPar2)
	case "redis":
		return ZoneMultiplexExcept(scw.ZoneFrPar3)
	default:
		return ZoneMultiplex
	}
}

func RegionMultiplexService(service string) schema.Multiplexer {
	switch service {
	case "tem":
		return RegionMultiplexSelective(scw.RegionFrPar)
	default:
		return RegionMultiplex
	}
}
