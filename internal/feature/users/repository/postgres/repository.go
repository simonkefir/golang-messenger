package users_repository_postgres

import (
	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

type UserRepository struct {
	pool core_postgres_pool.Pool
}

func NewUserRepository(pool core_postgres_pool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
