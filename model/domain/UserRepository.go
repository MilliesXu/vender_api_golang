package domain

type UserRepository interface {
	Register()
	Update()
	FindProfile()
	FindData()
	Verification()
	RequestChangePassword()
	ResetPassword()
}