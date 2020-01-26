package cndb

import (
	"awesomeProject7/gateway"
	"database/sql"
)

type cndb struct {
	db *sql.DB
}

func (c *cndb) BeginTx() (gateway.ContentTranslator, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	return &cndbTx{
		db: tx,
	}, nil
}

func (c *cndb) ToServiceGateway() gateway.ServiceGateway {
	return c
}

func (c *cndb) ToCategoryGateway() gateway.CategoryGateway {
	return c
}

func (c *cndb) ToContentGateway() gateway.ContentGateway {
	return c
}
