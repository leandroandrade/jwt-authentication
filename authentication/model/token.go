package model

import "time"

type Token struct {
	Token     string    `json:"token" form:"token"`
	CreatedAt time.Time `json:"-"`
}
