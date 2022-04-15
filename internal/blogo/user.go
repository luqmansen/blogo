package blogo

import "time"

type User struct {
	ID uint64 `json:"id"`

	Username string `json:"username"`
	Password string `json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
