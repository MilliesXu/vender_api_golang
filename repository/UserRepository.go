package repository

import (
	"context"
	"database/sql"
	"vender-api/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User
	Update(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User
	FindProfile(ctx context.Context, tx *sql.Tx) (*domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error)
	Verification(ctx context.Context, tx *sql.Tx, verificationCode string) (*domain.User, error)
	RequestChangePassword(ctx context.Context, tx *sql.Tx, email string)
	ResetPassword(ctx context.Context, tx *sql.Tx, password string, passwordResetCode string)
}
