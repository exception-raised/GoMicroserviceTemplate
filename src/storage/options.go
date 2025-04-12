package storage

import "github.com/exception-raised/microservices-template/src/storage/database"

type Option func(*Client)

func WithDatabase(db *database.Client) Option {
	return func(c *Client) {
		c.db = db
	}
}
