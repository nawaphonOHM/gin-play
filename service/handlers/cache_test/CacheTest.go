package cache_test

type CacheTest interface {
	Get(key string) (string, error)
	Set(key string, value string) (bool, error)
}
