package config

type Jwt struct {
	Exp     int64  `json:"exp" yaml:"exp"`
	Iss     string `json:"iss" yaml:"iss"`
	SignKey string `json:"signKey" yaml:"signKey"`
}
