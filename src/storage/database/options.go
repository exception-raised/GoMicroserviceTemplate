package database

import (
	"database/sql"

	"github.com/exception-raised/microservices-template/src/sqlc"
)

type Option func(*Client)

func WithDatabaseConnection(db *sql.DB) Option {
	return func(c *Client) {
		c.queries = sqlc.New(db)
	}
}
