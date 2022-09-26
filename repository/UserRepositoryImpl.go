package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"vender-api/helper"
	"vender-api/model/domain"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
}

func NewUserRepositry() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User {
	user.VerificationCode = uuid.New().String()
	user.PublicId = uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	helper.PanicIfError(err)

	user.Password = string(hashedPassword)

	SQL := "INSERT INTO user(public_id, firstname, lastname, email, password, verification_code, created_at, updated_at VALUES (?, ?, ?, ?, ?, ?, ?, ?))"

	result, err := tx.ExecContext(ctx, SQL, user.PublicId, user.Firstname, user.Lastname, user.Email, user.Password, user.VerificationCode, time.Now(), time.Now())

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User {
	userResult, err := repository.FindProfile(ctx, tx)

	userResult.Firstname = user.Firstname
	userResult.Lastname = user.Lastname
	userResult.Email = user.Email

	helper.PanicIfError(err)

	SQL := "UPDATE user SET firstname = ?, lastname = ?, email = ?, updated_at = ? where id = ?"
	_, err = tx.ExecContext(ctx, SQL, &userResult.Firstname, &userResult.Lastname, &userResult.Email, time.Now(), userResult.Id)

	helper.PanicIfError(err)

	return userResult
}

func (repository *UserRepositoryImpl) FindProfile(ctx context.Context, tx *sql.Tx) (*domain.User, error) {
	userId := ctx.Value("user_id").(int)

	SQL := "SELECT public_id, firstname, lastname, email, password, password_reset_code FROM user WHERE id = ?"

	rows, err := tx.QueryContext(ctx, SQL, userId)

	helper.PanicIfError(err)

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.PublicId, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.PasswordResetCode)

		helper.PanicIfError(err)

		return &user, nil
	}

	return &user, errors.New("user not found")
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	SQL := "SELECT id, public_id, firstname, lastname, email, verification_code, password_reset_code FROM user WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)

	helper.PanicIfError(err)

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.PublicId, &user.Firstname, &user.Lastname, &user.Email, &user.VerificationCode, &user.PasswordResetCode)

		helper.PanicIfError(err)

		return &user, nil
	}

	return &user, errors.New("user not found")

}

func (repository *UserRepositoryImpl) Verification(ctx context.Context, tx *sql.Tx, verificationCode string) (*domain.User, error) {
	userId := ctx.Value("user_id").(int)

	SQL := "SELECT firstname, lastname, email FROM user WHERE id = ?"

	rows, err := tx.QueryContext(ctx, SQL, userId)

	helper.PanicIfError(err)

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.Firstname, &user.Lastname, &user.Email)

		helper.PanicIfError(err)

		if user.VerificationCode != verificationCode {
			return &user, errors.New("user not found")
		}

		return &user, nil
	}

	return &user, errors.New("user not found")
}

func (repository *UserRepositoryImpl) RequestChangePassword(ctx context.Context, tx *sql.Tx, email string) {
	user, err := repository.FindByEmail(ctx, tx, email)

	helper.PanicIfError(err)

	user.PasswordResetCode = uuid.NewString()

	SQL := "UPDATE user SET password_reset_code = ?, updated_at = ? WHERE id = ?"

	_, err = tx.ExecContext(ctx, SQL, user.PasswordResetCode, time.Now(), user.Id)

	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) ResetPassword(ctx context.Context, tx *sql.Tx, password string, passwordResetCode string) {
	user, err := repository.FindProfile(ctx, tx)

	helper.PanicIfError(err)

	if user.PasswordResetCode != passwordResetCode {
		panic(errors.New("access denied"))
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	helper.PanicIfError(err)

	user.Password = string(newPassword)

	SQL := "UPDATE user SET password_reset_code = ?, password = ?, updated_at = ? WHERE id = ?"

	_, err = tx.ExecContext(ctx, SQL, nil, &user.Password, time.Now(), &user.Id)

	helper.PanicIfError(err)
}
