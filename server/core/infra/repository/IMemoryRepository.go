package repository

import "time"

type IMemoryRepository interface {
	Get(key string) *[]byte
	Set(key string, value []byte, expire time.Duration)
	Delete(key string)
	SetNX(key string, value []byte, expire time.Duration) bool
}
