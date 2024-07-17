package service_cache_test

type CacheTestService interface {
	Get(key string) string
	Set(key string, value string)
}
