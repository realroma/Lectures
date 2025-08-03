package buisnes

type Storage interface {
	GetOne()
	GetAll()
	Create()
	Edit(b Buisnes)
	Delete()
}
