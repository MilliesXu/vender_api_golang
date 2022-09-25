package repository

type ProductRepository interface {
	Create()
	Update()
	Delete()
	FindById()
	FindAll()
}
