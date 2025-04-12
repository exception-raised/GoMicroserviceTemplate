package config

import (
	"context"
	"database/sql"

	"github.com/exception-raised/microservices-template/src/storage/database"
)

func NewDatabaseStore(ctx context.Context, db *sql.DB) (*database.Client, error) {
	opts := []database.Option{
		database.WithDatabaseConnection(db),
	}

	return database.NewClient(opts...)
}
