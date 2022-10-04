package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"vender-api/helper"
	"vender-api/model/domain"
)

type MaterialRepositoryImpl struct{}

func NewMaterialRepository() MaterialRepository {
	return &MaterialRepositoryImpl{}
}

func (repository *MaterialRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, material *domain.Material) *domain.Material {
	userId := ctx.Value("userId").(int)

	SQL := "INSERT INTO material(name, description, uom, unit_price, user_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, material.Name, material.Description, material.Uom, material.UnitPrice, userId, time.Now(), time.Now())

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	material.Id = int(id)

	return material
}

func (repository *MaterialRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, material *domain.Material) *domain.Material {
	SQL := "UPDATE material SET name = ?, description = ?, uom = ?, unit_price = ?, updated_at = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, material.Name, material.Uom, material.UnitPrice, time.Now(), material.Id)

	helper.PanicIfError(err)

	return material
}

func (repository *MaterialRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, materialId int) {
	SQL := "DELETE FROM material WHERE id = ?"

	_, err := tx.ExecContext(ctx, SQL, materialId)

	helper.PanicIfError(err)
}

func (repository *MaterialRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, materialId int) (*domain.Material, error) {
	SQL := "SELECT m.id, m.name, m.description, m.uom, m.unit_price, m.created_at, m.updated_at, u.firstname, u.lastname FROM material INNER JOIN user u ON u.id = m.user_id WHERE m.id = ?"

	rows, err := tx.QueryContext(ctx, SQL, materialId)

	helper.PanicIfError(err)

	material := domain.Material{}
	material.User = map[string]string{}

	if rows.Next() {
		err := rows.Scan(material.Id, material.Name, material.Description, material.Uom, material.UnitPrice, material.CreatedAt, material.UpdatedAt, material.User["firstname"], material.User["lastname"])

		helper.PanicIfError(err)

		return &material, nil
	}

	return &material, errors.New("not found")
}

func (repository *MaterialRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []*domain.Material {
	SQL := "SELECT m.id, m.name, m.description, m.uom, m.unit_price, m.created_at, m.updated_at, u.firstname, u.lastname FROM material INNER JOIN user u ON u.id = m.user_id"

	rows, err := tx.QueryContext(ctx, SQL)

	helper.PanicIfError(err)

	materials := []*domain.Material{}

	for rows.Next() {
		material := domain.Material{}
		material.User = map[string]string{}

		err := rows.Scan(material.Id, material.Name, material.Description, material.Uom, material.UnitPrice, material.CreatedAt, material.UpdatedAt, material.User["firstname"], material.User["lastname"])

		helper.PanicIfError(err)

		materials = append(materials, &material)
	}

	return materials
}
