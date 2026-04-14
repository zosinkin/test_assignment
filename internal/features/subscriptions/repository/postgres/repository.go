package subscriptions_postgres_repository

import core_postgres_pool "github.com/zosinkin/test_assignment.git/internal/core/repository/postgres/pool"


type SubRepository struct {
	pool core_postgres_pool.Pool
}


func NewSubRepository(
	pool core_postgres_pool.Pool,
) *SubRepository {
	return &SubRepository{
		pool: pool,
	}
}