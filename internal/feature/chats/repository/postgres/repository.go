package chats_repository_postgres

import core_postgres_pool "github.com/simonkefir/golang-messenger/internal/core/repository/postgres/pool"

type ChatRepository struct {
	pool core_postgres_pool.Pool
}

func NewChatRepository(pool core_postgres_pool.Pool) *ChatRepository {
	return &ChatRepository{
		pool: pool,
	}
}
