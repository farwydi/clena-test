package usecase

import (
	"awesomeProject7/domain"
	"awesomeProject7/gateway"
)

type ServiceUseCase interface {
	EnableTx() (gateway.ContentTranslator, error)
	CopyService(in *domain.Service) (out *domain.Service, err error)
}

type ContentUseCase interface {
	CopyContent(*domain.Content) (*domain.Content, error)
}

type CategoryUseCase interface {
	CopyCategory(*domain.Category) (*domain.Category, error)
}
