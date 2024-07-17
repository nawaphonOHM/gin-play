package cache_test

type CacheTestHandler interface {
	Get(key string) (string, error)
	Set(key string, value string) (bool, error)
}
