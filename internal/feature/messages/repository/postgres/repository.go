package messages_repository_postgres

import (
	core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"
)

type MsgRepository struct {
	pool core_postgres_pool.Pool
}

func NewMsgRepository(pool core_postgres_pool.Pool) *MsgRepository {
	return &MsgRepository{
		pool: pool,
	}
}
