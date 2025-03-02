package utils

import (
	"errors"
	"neocex/v2/config"
	"neocex/v2/global"
	sysReq "neocex/v2/internal/models/global/req"
)

type JWT struct {
	SigningKey  []byte
	BufferTime  string
	ExpiredTime string
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT(config *config.JWT) *JWT {
	if config == nil {
		//Acquire jwt config from default config
		config = &global.GVA_CONF.JWT
	}
	return &JWT{
		SigningKey:  []byte(config.SigningKey),
		BufferTime:  config.BufferTime,
		ExpiredTime: config.ExpiredTime,
	}
}

func (j *JWT) ParseToken(token string) (*sysReq.CustomerClaims, error) {
	// token, err := jwt.Parse()
	//Hard code for new, and development for later
	return new(sysReq.CustomerClaims), nil
}
