package domain

type ProductRepository interface {
	Create()
	Update()
	Delete()
	FindById()
	FindAll()
}
