package redisconfig

// Config configures a redis cache
type RedisConfig struct {
	address string
}

//GetAddress
func (cfg *RedisConfig) GetAddress() string {
	return cfg.address
}

// NewRedisConfig returns a redis config based on environment variables
func NewRedisConfig(redisEnvKey string) RedisConfig {
	addr := redisEnvKey
	return RedisConfig{addr}
}


//TODO: Implement the redis cache storage to save and retrieve the game state
