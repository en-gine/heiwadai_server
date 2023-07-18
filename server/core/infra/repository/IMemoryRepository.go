package repository

type IMemoryRepository interface {
	Get(key string) (string, error)
	Save(key string, value string) error
	Delete(key string) error
}
