package repository

import (
	"context"
	"database/sql"
	"vender-api/helper"
	"vender-api/model/domain"
)

type SessionRepositoryImpl struct {
}

func NewSessionRepository() SessionRepository {
	return &SessionRepositoryImpl{}
}

func (repository *SessionRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, session *domain.Session) *domain.Session {
	userId := ctx.Value("user_id").(int)

	SQL := "INSERT INTO session(user_id, valid, created_at, updated_at) VALUES(?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, userId, session.Valid, session.CreatedAt, session.UpdatedAt)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	session.Id = int(id)

	return session
}

func (repository *SessionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, session *domain.Session) {
	sessionId := ctx.Value("session_id").(int)

	SQL := "UPDATE session SET valid = ?, updated_at = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, SQL, session.Valid, session.UpdatedAt, sessionId)

	helper.PanicIfError(err)
}
