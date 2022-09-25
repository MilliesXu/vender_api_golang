package domain

import "time"

type Product struct {
	Id            int
	Name          string
	Description   string
	MaterialLines []ProductMaterialLine
	User          map[string]string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
