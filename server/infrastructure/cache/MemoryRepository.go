package cache

import (
	"errors"
	"time"

	"server/core/infra/repository"

	"github.com/patrickmn/go-cache"
)

var _ repository.IMemoryRepository = &MemoryRepository{}

type MemoryRepository struct {
	db *cache.Cache
}

func NewMemoryRepository() (*MemoryRepository, error) {
	db := cache.New(60*time.Minute, 60*time.Minute)
	return &MemoryRepository{
		db: db,
	}, nil
}

func (mr *MemoryRepository) Get(key string) *[]byte {
	cache, found := mr.db.Get(key)
	if !found || cache == nil {
		return nil
	}

	data, ok := cache.([]byte)
	if !ok {
		return nil
	}
	return &data
}

func (mr *MemoryRepository) Set(key string, value []byte, expire time.Duration) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	if value == nil || len(value) == 0 {
		return errors.New("value cannot be nil or empty")
	}

	mr.db.Set(key, value, expire)

	return nil
}

func (mr *MemoryRepository) Delete(key string) {
	mr.db.Delete(key)
}
