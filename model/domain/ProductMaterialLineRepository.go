package domain

type ProductMaterialLineRepository interface {
	Create()
	Update()
	Delete()
	FindById()
}
