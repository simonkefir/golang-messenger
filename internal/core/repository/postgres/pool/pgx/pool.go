package core_pgx_pool

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

type Pool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

type Tx struct {
	tx pgx.Tx
}

func NewPool(
	ctx context.Context,
	config Config,
) (*Pool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	pgxconfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("parse pgxconfig: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)
	if err != nil {
		return nil, fmt.Errorf("create pgxpool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgxpool ping: %w", err)
	}

	return &Pool{
		Pool:      pool,
		opTimeout: config.Timeout,
	}, nil
}

func (p *Pool) Query(
	ctx context.Context,
	sql string,
	args ...any,
) (core_postgres_pool.Rows, error) {
	rows, err := p.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return pgxRows{rows}, nil
}

func (p *Pool) QueryRow(
	ctx context.Context,
	sql string,
	args ...any,
) core_postgres_pool.Row {
	row := p.Pool.QueryRow(ctx, sql, args...)

	return pgxRow{row}
}

func (p *Pool) Exec(ctx context.Context,
	sql string,
	arguments ...any,
) (core_postgres_pool.CommandTag, error) {
	tag, err := p.Pool.Exec(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}

	return pgxCommandTag{tag}, nil
}

func (p *Pool) OpTimeOut() time.Duration {
	return p.opTimeout
}

func (p *Pool) Begin(ctx context.Context) (core_postgres_pool.Tx, error) {
	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &Tx{
		tx: tx,
	}, nil
}

func (t *Tx) Query(ctx context.Context, sql string, args ...any) (core_postgres_pool.Rows, error) {
	rows, err := t.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &pgxRows{rows}, nil
}

func (t *Tx) QueryRow(ctx context.Context, sql string, args ...any) core_postgres_pool.Row {
	return &pgxRow{t.tx.QueryRow(ctx, sql, args...)}
}

func (t *Tx) Exec(ctx context.Context, sql string, args ...any) (core_postgres_pool.CommandTag, error) {
	tag, err := t.tx.Exec(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	return &pgxCommandTag{tag}, nil
}

func (t *Tx) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *Tx) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}
