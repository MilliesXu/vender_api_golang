package domain

type MaterialRepository interface {
	Create()
	Update()
	Delete()
	FindById()
	FindAll()
}
