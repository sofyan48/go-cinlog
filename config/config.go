package config

// Configs ...
type Configs struct {
	Service string
	URL     string
}

// NewConfig ...
func NewConfig() *Configs {
	return &Configs{}
}

// SetClient ...
func (cfg *Configs) SetClient(url, svc string) *Configs {
	cfg.Service = svc
	cfg.URL = url
	return cfg
}
