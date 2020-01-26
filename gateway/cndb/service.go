package cndb

import (
	"awesomeProject7/domain"
	"awesomeProject7/gateway"
	"database/sql"
)

func NewServiceGateway(db *sql.DB) gateway.ServiceGateway {
	return &cndb{
		db: db,
	}
}

func (c *cndbTx) CreateService(name string) (*domain.Service, error) {
	panic("implement me")
}

func (c *cndb) CreateService(name string) (*domain.Service, error) {
	panic("implement me")
}

func (c *cndbTx) UpdateService(id int, name string) (*domain.Service, error) {
	panic("implement me")
}

func (c *cndb) UpdateService(id int, name string) (*domain.Service, error) {
	panic("implement me")
}

func (c *cndbTx) DeleteService(id int) error {
	panic("implement me")
}

func (c *cndb) DeleteService(id int) error {
	panic("implement me")
}

func (c *cndbTx) CountService() (int, error) {
	panic("implement me")
}

func (c *cndb) CountService() (int, error) {
	panic("implement me")
}
