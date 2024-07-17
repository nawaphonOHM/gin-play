package cache_test

type CacheTestService interface {
	Get(key string) string
}
