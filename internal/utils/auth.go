package utils

import (
	"neocex/v2/global"
	sysReq "neocex/v2/internal/models/global/req"
	"net"
	"strconv"
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
	//Add cookie to source of network
	host, _, err := net.SplitHostPort(c.Hostname())
	if err != nil {
		host = c.Hostname()
	}
	if net.ParseIP(host) != nil {
		c.Request().Header.SetCookie("token", token)
		c.Request().Header.SetCookie("Path:", "/")
		c.Request().Header.SetCookie("Domain:", host)
		c.Request().Header.SetCookie("Max-Age:", strconv.Itoa(maxAge*60))
	} else {
		c.Request().Header.SetCookie("token", token)
		c.Request().Header.SetCookie("Path:", "/")
		c.Request().Header.SetCookie("Max-Age:", strconv.Itoa(maxAge*60))
	}
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
