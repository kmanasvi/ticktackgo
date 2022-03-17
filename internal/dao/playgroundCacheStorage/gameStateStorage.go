package gameStorage

import redisconfig "github.com/tictacgo/config/redis"

type cacheStorage struct {
	redisconfig.RedisConfig
}

//constructor for storage
func NewcacheStorage(redisConfig redisconfig.RedisConfig) *cacheStorage {
	return &cacheStorage{redisConfig}
}

//TODO-  Implement redis cache storage to save the game state
