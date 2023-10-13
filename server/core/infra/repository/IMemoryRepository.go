package repository

import "time"

type IMemoryRepository interface {
	Get(key string) *[]byte
	Set(key string, value []byte, expire time.Duration) error
	Delete(key string)
}
