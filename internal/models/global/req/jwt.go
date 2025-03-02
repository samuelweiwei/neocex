package req

import (
	"github.com/golang-jwt/jwt/v4"
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
	Usernmae       string    `json:"username"`
	PreferableName string    `json:"preferable_name"`
	AuthorityId    uint      `json:"authority_id"`
	FrontendUserId uint      `json:"frontend_user_id"`
	UserType       uint      `json:"user_type"`
}
