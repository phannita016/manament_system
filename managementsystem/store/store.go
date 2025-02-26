package store

type Cache[T any] interface {
	Get(key string) (T, bool)
	Set(key string, t T) bool
	Delete(key string) bool
	Keys() []string
	Values() []T
}
