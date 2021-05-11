package cache

import "web/config"

func NewCacheManager(cfg config.CacheConfig) (Manager, error) {
	switch cfg.Driver {
	case "redis":
		return newRedisCache(cfg)
	default:
		return newRedisCache(cfg)
	}
}
