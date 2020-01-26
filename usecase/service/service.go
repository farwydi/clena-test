package service

import (
	"awesomeProject7/domain"
	"awesomeProject7/gateway"
	"awesomeProject7/usecase"
)

func New(gw gateway.ServiceGateway) usecase.ServiceUseCase {
	return &serviceUseCase{
		gw: gw,
	}
}

type serviceUseCase struct {
	gw gateway.ServiceGateway
}

func (s *serviceUseCase) EnableTx() (gateway.ContentTranslator, error) {
	return s.gw.BeginTx()
}

func CopyServiceTx(gw gateway.ServiceGateway, in *domain.Service) (out *domain.Service, err error) {
	return gw.CreateService(in.Name)
}

func (s *serviceUseCase) CopyService(in *domain.Service) (out *domain.Service, err error) {
	return CopyServiceTx(s.gw, in)
}
