package domain

import "time"

type Session struct {
	Id         int
	User       map[string]string
	Valid      bool
	Created_at time.Time
	Updated_at time.Time
}
