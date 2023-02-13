package client

import "github.com/scaleway/scaleway-sdk-go/scw"

type Spec struct {
	OrgIDs  []string     `json:"org_ids,omitempty"`
	Regions []scw.Region `json:"regions,omitempty"`
	Zones   []scw.Zone   `json:"zones,omitempty"`

	// Optional
	Timeout int64 `json:"timeout_secs,omitempty"`
}

func (s Spec) Validate() error {
	return nil
}

func (s *Spec) SetDefaults() {
	if len(s.Regions) == 0 {
		s.Regions = scw.AllRegions
	}
	if len(s.Zones) == 0 {
		s.Zones = scw.AllZones
	}

	if s.Timeout < 1 {
		s.Timeout = 10
	}
}
