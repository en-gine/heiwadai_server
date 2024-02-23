package redis

import (
	"context"
	"time"

	"server/core/infra/repository"
	"server/infrastructure/env"

	"github.com/redis/go-redis/v9"
)

var (
	_   repository.IMemoryRepository = &MemoryRepository{}
	ctx                              = context.Background()
)

type MemoryRepository struct {
	db *redis.Client
}

func NewMemoryRepository() *MemoryRepository {
	db := redis.NewClient(&redis.Options{
		Addr:     env.GetEnv(env.RedisHost) + ":" + env.GetEnv(env.RedisPort),
		Password: env.GetEnv(env.RedisPass),
		Username: env.GetEnv(env.RedisUser),
		DB:       0,
	})
	return &MemoryRepository{
		db: db,
	}
}

func (mr *MemoryRepository) Get(key string) *[]byte {
	cache := mr.db.Get(ctx, key)
	if cache.Err() != nil {
		return nil
	}
	data, _ := cache.Bytes()

	return &data
}

func (mr *MemoryRepository) Set(key string, value []byte, expire time.Duration) {
	mr.db.Set(ctx, key, value, expire)
}

func (mr *MemoryRepository) Delete(key string) {
	mr.db.Del(ctx, key)
}
