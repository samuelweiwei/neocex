package config

type JWT struct {
	SigningKey  string `mapstructure:"signing_key" json:"signing_key" yaml:"signing_key"`
	BufferTime  string `mapstructure:"buffer_time" json:"buffer_time" yaml:"buffer_time"`
	ExpiredTime string `mapstructure:"expired_time" json:"expired_time" yaml:"expired_time"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
