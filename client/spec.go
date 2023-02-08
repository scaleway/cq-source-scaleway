package client

type Spec struct {
	// Optional
	Timeout int64 `json:"timeout_secs,omitempty"`
}

func (s Spec) Validate() error {
	return nil
}

func (s *Spec) SetDefaults() {
	if s.Timeout < 1 {
		s.Timeout = 10
	}
}
