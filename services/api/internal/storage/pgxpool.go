package storage

import (
	"context"
	"fmt"
	"net/url"

	"github.com/DavelPurov777/microblog/configs/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPgxPool(ctx context.Context, cfg config.ConfigPgxpool, poolCfg config.PoolSettings) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		url.PathEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	pgxCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	if poolCfg.MaxConns > 0 {
		pgxCfg.MaxConns = poolCfg.MaxConns
	}
	pgxCfg.MinConns = poolCfg.MinConns
	if poolCfg.MaxConnLifeTime > 0 {
		pgxCfg.MaxConnLifetime = poolCfg.MaxConnLifeTime
	}
	if poolCfg.MaxConnIdleTime > 0 {
		pgxCfg.MaxConnIdleTime = poolCfg.MaxConnIdleTime
	}
	if poolCfg.HealthCheckPeriod > 0 {
		pgxCfg.HealthCheckPeriod = poolCfg.HealthCheckPeriod
	}
	pool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

func NewSQLXFromPgxPool(pool *pgxpool.Pool) *sqlx.DB {
	stdDB := stdlib.OpenDBFromPool(pool)
	return sqlx.NewDb(stdDB, "pgx")
}
