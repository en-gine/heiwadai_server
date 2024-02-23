package cache

import (
	"time"

	"server/core/infra/repository"

	"github.com/patrickmn/go-cache"
)

var _ repository.IMemoryRepository = &MemoryRepository{}

type MemoryRepository struct {
	db *cache.Cache
}

func NewMemoryRepository() *MemoryRepository {
	db := cache.New(60*time.Minute, 60*time.Minute)
	return &MemoryRepository{
		db: db,
	}
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

func (mr *MemoryRepository) Set(key string, value []byte, expire time.Duration) {
	mr.db.Set(key, value, expire)
}

func (mr *MemoryRepository) Delete(key string) {
	mr.db.Delete(key)
}
