package repository

import (
	"context"
	"database/sql"
	"vender-api/model/domain"
)

type SessionRepository interface {
	Create(ctx context.Context, tx *sql.Tx, session *domain.Session) *domain.Session
	Update(ctx context.Context, tx *sql.Tx, session *domain.Session)
}
