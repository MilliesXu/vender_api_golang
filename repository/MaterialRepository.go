package repository

type MaterialRepository interface {
	Create()
	Update()
	Delete()
	FindById()
	FindAll()
}
