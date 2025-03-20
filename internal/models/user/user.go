package user

import "time"

type User struct {
	Address    string
	CreateTime time.Time
	UpdateTime time.Time
}
