package config

type Server struct {
	JWT     JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	JWTUser JWT `mapstructure:"jwt_user" json:"jwt_user" yaml:"jwt_user"`
}
