package cndb

import (
	"awesomeProject7/domain"
	"awesomeProject7/gateway"
)

func NewCategoryGateway() gateway.CategoryGateway {
	return &cndb{}
}

func (c *cndbTx) CreateCategory(name string) (*domain.Category, error) {
	_, err := c.db.Exec("", name)
	if err != nil {
		return nil, err
	}

	return &domain.Category{}, nil
}

func (c *cndb) CreateCategory(name string) (*domain.Category, error) {
	_, err := c.db.Exec("", name)
	if err != nil {
		return nil, err
	}

	return &domain.Category{}, nil
}

func (c *cndbTx) UpdateCategory(id int, name string) (*domain.Category, error) {
	panic("implement me")
}

func (c *cndb) UpdateCategory(id int, name string) (*domain.Category, error) {
	panic("implement me")
}

func (c *cndbTx) DeleteCategory(id int) error {
	panic("implement me")
}

func (c *cndb) DeleteCategory(id int) error {
	panic("implement me")
}

func (c *cndbTx) CountCategory() (int, error) {
	panic("implement me")
}

func (c *cndb) CountCategory() (int, error) {
	panic("implement me")
}
