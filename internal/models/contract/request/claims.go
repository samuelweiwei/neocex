package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type CustomerClaims struct {
	BaseClaims
	BufferTime int64 `json:"buffer_time"`
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID           uuid.UUID `json:"uuid"`
	ID             uint      `json:"id"`
	Username       string    `json:"username"`
	PreferableName string    `json:"preferable_name"`
	AuthorityId    string    `json:"authority_id"`
	FrontUserId    string    `json:"front_user_id"`
	UserType       string    `json:"user_type"`
}
