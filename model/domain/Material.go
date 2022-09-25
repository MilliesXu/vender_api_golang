package domain

import "time"

type Material struct {
	Id          int
	Name        string
	Description string
	Uom         string
	UnitPrice   float64
	User        map[string]string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
