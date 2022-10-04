package repository

import (
	"context"
	"database/sql"
	"vender-api/model/domain"
)

type MaterialRepository interface {
	Create(ctx context.Context, tx *sql.Tx, material *domain.Material) *domain.Material
	Update(ctx context.Context, tx *sql.Tx, material *domain.Material) *domain.Material
	Delete(ctx context.Context, tx *sql.Tx, materialId int)
	FindById(ctx context.Context, tx *sql.Tx, materialId int) (*domain.Material, error)
	FindAll(ctx context.Context, tx *sql.Tx) []*domain.Material
}
