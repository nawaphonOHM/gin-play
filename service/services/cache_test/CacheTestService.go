package service_cache_test

type CacheTestService interface {
	Get(key string) string
}
