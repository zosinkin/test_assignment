package core_pgx_pool


import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	core_postgres_pool "github.com/zosinkin/test_assignment.git/internal/core/repository/postgres/pool"
)

// Pool — обёртка над pgxpool.Pool с добавлением конфигурации таймаута операций.
type Pool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}


// NewPool создаёт и инициализирует новый пул соединений с PostgreSQL.
// Формирует connection string, парсит конфигурацию, создаёт пул и проверяет соединение через Ping.
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
		Pool: pool,
		opTimeout: config.Timeout,
	}, nil
}


// Query выполняет SQL-запрос, возвращающий несколько строк.
// Оборачивает pgxpool.Query и возвращает интерфейс Rows.
func (p *Pool) Query(
	ctx context.Context, 
	sql string, args ...any,
) (core_postgres_pool.Rows, error) {
	rows, err := p.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
	
}


// QueryRow выполняет SQL-запрос, возвращающий одну строку.
// Оборачивает pgxpool.QueryRow.
func(p *Pool) QueryRow(
	ctx context.Context,
	sql string, 
	args ...any,
) core_postgres_pool.Row {
	row := p.Pool.QueryRow(ctx, sql, args...)
	return row
}


// Exec выполняет SQL-запрос без возврата строк (INSERT, UPDATE, DELETE).
// Возвращает CommandTag с информацией о выполненной операции.
func (p *Pool) Exec(
	ctx context.Context, 
	sql string, 
	arguments ...any,
) (core_postgres_pool.CommandTag, error) {
	tag, err := p.Pool.Exec(ctx, sql, arguments...)
	if err != nil {
		return nil, err
	}

	return pgxComandTag{tag}, nil
}


// OpTimeout возвращает таймаут операций с БД,
// который используется для ограничения времени выполнения запросов
func (p *Pool) OpTimeout() time.Duration {
	return p.opTimeout
}