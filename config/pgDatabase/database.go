package pgDatabase

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	index "github.com/tictacgo/config/environment"
)

func SetDBConnection() (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     index.EnvironmentConfig.DB_ADDR,
		User:     index.EnvironmentConfig.DB_USER,
		Password: index.EnvironmentConfig.DB_PASSWORD,
		Database: index.EnvironmentConfig.DB_DATABASE,
	})
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "Error in pgDatabase config while testing ping for database connection")
	}
	_, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		return nil, errors.Wrap(err, "Error in pgDatabase config while testing a sample query in the data base")
	}
	return db, nil
}
