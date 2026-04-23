package utils

import (
	"errors"
	"neocex/v2/config"
	req "neocex/v2/internal/models/global/req"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Sentinel errors returned by ParseToken.
var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("token is malformed")
	TokenInvalid     = errors.New("token is invalid")
)

// JWT holds signing configuration.
type JWT struct {
	signingKey  []byte
	expiredTime time.Duration
	bufferTime  time.Duration
	issuer      string
}

// NewJWT builds a JWT helper from a config.JWT value.
func NewJWT(cfg *config.JWT) *JWT {
	expired, err := time.ParseDuration(cfg.ExpiredTime)
	if err != nil {
		expired = 24 * time.Hour
	}
	buffer, err := time.ParseDuration(cfg.BufferTime)
	if err != nil {
		buffer = time.Hour
	}
	return &JWT{
		signingKey:  []byte(cfg.SigningKey),
		expiredTime: expired,
		bufferTime:  buffer,
		issuer:      cfg.Issuer,
	}
}

// GenerateToken creates and signs a JWT containing the given BaseClaims.
func (j *JWT) GenerateToken(baseClaims req.BaseClaims) (string, error) {
	now := time.Now()
	claims := req.CustomerClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(j.bufferTime.Seconds()),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.expiredTime)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}

// ParseToken validates tokenStr and returns the embedded CustomerClaims.
func (j *JWT) ParseToken(tokenStr string) (*req.CustomerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &req.CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, TokenInvalid
		}
		return j.signingKey, nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, TokenExpired
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, TokenNotValidYet
		case errors.Is(err, jwt.ErrTokenMalformed):
			return nil, TokenMalformed
		default:
			return nil, TokenInvalid
		}
	}
	if claims, ok := token.Claims.(*req.CustomerClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken re-signs a token if it is still within the buffer window past expiry.
func (j *JWT) RefreshToken(tokenStr string) (string, error) {
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	token, _, err := parser.ParseUnverified(tokenStr, &req.CustomerClaims{})
	if err != nil {
		return "", TokenMalformed
	}
	claims, ok := token.Claims.(*req.CustomerClaims)
	if !ok {
		return "", TokenInvalid
	}
	if claims.ExpiresAt != nil && time.Since(claims.ExpiresAt.Time) > j.bufferTime {
		return "", TokenExpired
	}
	return j.GenerateToken(claims.BaseClaims)
}
