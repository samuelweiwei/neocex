package utils

import (
	"neocex/v2/global"
	sysReq "neocex/v2/internal/models/global/req"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetClaims(f *fiber.Ctx) (*sysReq.CustomerClaims, error) {
	token := GetToken(f)
	j := NewJWT(&global.GVA_CONF.JWT)
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, err
}

func GetToken(f *fiber.Ctx) string {
	token := f.Cookies("token", "")
	if token == "" {
		j := NewJWT(&global.GVA_CONF.JWT)
		token = f.Get("token")
		claims, err := j.ParseToken(token)
		if err != nil {
			return ""
		}
		SetToken(f, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	}
	return token
}

func SetToken(c *fiber.Ctx, token string, maxAge int) {
	host, _, err := net.SplitHostPort(c.Hostname())
	if err != nil {
		host = c.Hostname()
	}
	cookie := &fiber.Cookie{
		Name:   "token",
		Value:  token,
		Path:   "/",
		MaxAge: maxAge * 60,
	}
	if net.ParseIP(host) == nil {
		// Only set Domain for named hosts, not raw IPs.
		cookie.Domain = host
	}
	c.Cookie(cookie)
}

func GetUserID(c *fiber.Ctx) uint {
	claims, exists := c.Locals("claims").(*sysReq.CustomerClaims)
	if !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		return claims.BaseClaims.ID
	}
}
