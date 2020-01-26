package gateway

import "awesomeProject7/domain"

type ServiceGateway interface {
	ContentTranslator

	CreateService(name string) (*domain.Service, error)
	//CreateServiceTx(tx Transaction, name string) (*domain.Service, error)
	UpdateService(id int, name string) (*domain.Service, error)
	//UpdateServiceTx(tx Transaction, id int, name string) (*domain.Service, error)
	DeleteService(id int) error
	//DeleteServiceTx(tx Transaction, id int) error
	CountService() (int, error)
}

type ContentGateway interface {
	ContentTranslator

	CreateContent(name string) (*domain.Content, error)
	//CreateContentTx(tx Transaction, name string) (*domain.Content, error)
	UpdateContent(id int, name string) (*domain.Content, error)
	//UpdateContentTx(tx Transaction, id int, name string) (*domain.Content, error)
	DeleteContent(id int) error
	//DeleteContentTx(tx Transaction, id int) error
	CountContent() (int, error)
}

type CategoryGateway interface {
	ContentTranslator

	CreateCategory(name string) (*domain.Category, error)
	//CreateCategoryTx(tx Transaction, name string) (*domain.Category, error)
	UpdateCategory(id int, name string) (*domain.Category, error)
	//UpdateCategoryTx(tx Transaction, id int, name string) (*domain.Category, error)
	DeleteCategory(id int) error
	//DeleteCategoryTx(tx Transaction, id int) error
	CountCategory() (int, error)
}
