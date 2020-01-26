package cndb

import (
	"awesomeProject7/gateway"
	"database/sql"
	"errors"
)

type cndbTx struct {
	db *sql.Tx
}

func (c *cndbTx) BeginTx() (gateway.ContentTranslator, error) {
	return nil, errors.New("transaction of transaction doesnt support")
}

func (c *cndbTx) ToServiceGateway() gateway.ServiceGateway {
	return c
}

func (c *cndbTx) ToCategoryGateway() gateway.CategoryGateway {
	return c
}

func (c *cndbTx) ToContentGateway() gateway.ContentGateway {
	return c
}
