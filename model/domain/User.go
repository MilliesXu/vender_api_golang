package domain

import "time"

type User struct {
	Id                int
	PublicId          string
	Firstname         string
	Lastname          string
	Email             string
	Password          string
	VerificationCode  string
	PasswordResetCode string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
