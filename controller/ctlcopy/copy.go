package ctlcopy

import (
	"awesomeProject7/gateway/cndb"
	"awesomeProject7/usecase"
	"awesomeProject7/usecase/service"
	"runtime"
)

type Sig func(id int) error

func CtlCopy(suc usecase.ServiceUseCase, ctuc usecase.CategoryUseCase, cnuc usecase.ContentUseCase) Sig {
	return func(id int) error {
		contentTx, err := suc.EnableTx()
		if err != nil {
			return err
		}

		s, err := service.CopyServiceTx(contentTx.ToServiceGateway(), nil)
		if err != nil {
			return err
		}

		return nil
	}
}
