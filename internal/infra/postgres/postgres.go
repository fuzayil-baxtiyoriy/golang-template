package postgres

import (
	"context"
	"fmt"
	"golang-template/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func NewConnPool(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "failed to connect to postgres")
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "failed to ping postgres")
	}

	return pool, nil
}
