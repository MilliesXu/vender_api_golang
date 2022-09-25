package domain

import "time"

type ProductMaterialLine struct {
	Id        int
	Material  Material
	Quantity  float64
	User      map[string]string
	CreatedAt time.Time
	UpdateAt  time.Time
}
