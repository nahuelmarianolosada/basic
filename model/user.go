package model

import "time"

type User struct {
	UID       int64     `json:"id"`
	Username string    `json:"username"`
	Created  time.Time `json:"created"`
}
