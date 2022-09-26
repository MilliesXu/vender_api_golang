package domain

import "time"

type Session struct {
	Id        int
	User      map[string]string
	Valid     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
