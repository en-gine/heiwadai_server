package repository

import "time"

type IMemoryRepository interface {
	Get(key string) (*[]byte, error)
	Set(key string, value []byte, expire time.Duration) error
	Delete(key string) error
}
