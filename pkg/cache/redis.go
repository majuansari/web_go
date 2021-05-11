package cache

import (
	ekocache "github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/go-redis/redis/v7"
	"web/config"
)

type redisManager struct {
	cache *ekocache.Cache
}

type Manager interface {
	Get(string, callBack, *store.Options) (interface{}, error)
	Delete(string) error
	Flush() error
	Take(string) (interface{}, error)
}

func newRedisCache(cfg config.CacheConfig) (*redisManager, error) {
	redisCache := initRedisCache(cfg)
	err := redisCache.Set("_init_", "test", &store.Options{})
	if err != nil {
		return nil, err
	}
	return &redisManager{redisCache}, nil
}

func initRedisCache(cfg config.CacheConfig) *ekocache.Cache {
	redisStore := store.NewRedis(redis.NewClient(&redis.Options{
		Addr: cfg.Host + ":" + cfg.Port,
	}), nil)

	return ekocache.New(redisStore)
}

type callBack func() (interface{}, error)

func (rm redisManager) Get(key string, cb callBack, options *store.Options) (interface{}, error) {
	//return rm.loadData(key, cb, options)

	value, err := rm.cache.Get(key)
	switch {
	case err == redis.Nil:
		return rm.loadData(key, cb, options)
	case err != nil:
		return nil, err
	case value != "":
		return value, nil
	}
	return nil, nil
}

func (rm redisManager) Take(key string) (interface{}, error) {
	return rm.cache.Get(key)
}
func (rm redisManager) Delete(key string) error {
	return rm.cache.Delete(key)
}

func (rm redisManager) Flush() error {
	return rm.cache.Clear()
}

func (rm redisManager) loadData(key string, cb callBack, options *store.Options) (interface{}, error) {
	value, err := cb()
	if err == nil {
		err := rm.cache.Set(key, value, options)
		if err != nil {
			//Log error
		}
		return value, nil
	}
	return nil, err
}
