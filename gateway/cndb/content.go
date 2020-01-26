package cndb

import (
	"awesomeProject7/domain"
	"awesomeProject7/gateway"
)

func NewContentGateway() gateway.ContentGateway {
	return &cndb{}
}

func (c *cndbTx) CreateContent(name string) (*domain.Content, error) {
	panic("implement me")
}

func (c *cndb) CreateContent(name string) (*domain.Content, error) {
	panic("implement me")
}

func (c *cndbTx) UpdateContent(id int, name string) (*domain.Content, error) {
	panic("implement me")
}

func (c *cndb) UpdateContent(id int, name string) (*domain.Content, error) {
	panic("implement me")
}

func (c *cndbTx) DeleteContent(id int) error {
	panic("implement me")
}

func (c *cndb) DeleteContent(id int) error {
	panic("implement me")
}

func (c *cndbTx) CountContent() (int, error) {
	panic("implement me")
}

func (c *cndb) CountContent() (int, error) {
	panic("implement me")
}
