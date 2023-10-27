package client

import "github.com/scaleway/scaleway-sdk-go/scw"

type Spec struct {
	Regions []scw.Region `json:"regions,omitempty"`
	Zones   []scw.Zone   `json:"zones,omitempty"`

	// Optional
	Timeout     int64 `json:"timeout_secs,omitempty"`
	Concurrency int   `json:"concurrency,omitempty"`
}

func (Spec) Validate() error {
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
	if s.Concurrency < 1 {
		s.Concurrency = 1000
	}
}
