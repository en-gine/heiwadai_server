package repository

import (
	"context"
	"os"
	"server/core/infra/repository"

	"github.com/redis/go-redis/v9"
)

var _ repository.IMemoryRepository = &MemoryRepository{}

type MemoryRepository struct {
	rdb *redis.Client
}

func NewMemoryRepository() (*MemoryRepository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + "" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	err := rdb.Ping(context.Background()).Err()

	if err != nil {
		return nil, err
	}

	return &MemoryRepository{
		rdb: rdb,
	}, nil
}

func (mr *MemoryRepository) Get(key string) (string, error) {
	value, err := mr.rdb.Get(context.Background(), key).Result()
	return value, err
}

func (mr *MemoryRepository) Save(key string, value string) error {
	err := mr.rdb.Set(context.Background(), key, value, 0).Err()
	return err
}

func (mr *MemoryRepository) Delete(key string) error {
	err := mr.rdb.Del(context.Background(), key).Err()
	return err
}
