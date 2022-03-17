package pubsub

import (
	"log"
	"time"

	redisgo "github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	redis "github.com/tictacgo/config/redis"
)

const (
	network = "tcp"
)

type RedisClient struct {
	pool        *redisgo.Pool
	channelport string
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func NewRedisClient(cfg redis.RedisConfig, channelport string) (*RedisClient, error) {
	pool := &redisgo.Pool{
		MaxIdle:     10,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redisgo.Conn, error) {
			c, err := redisgo.Dial(network, cfg.GetAddress())
			if err != nil {
				return nil, errors.Wrap(err, "Error while dialing to redis client")
			}
			return c, err
		},
	}
	err := Ping(pool.Get())
	if err != nil {
		return nil, errors.Wrap(err, "Error while pinging test for redis client connection ")
	}
	log.Print("Redis initiated for ", cfg.GetAddress())
	return &RedisClient{pool, channelport}, nil
}

func (r *RedisClient) ClosePool() {
	r.pool.Close()
}

func Ping(c redisgo.Conn) error {
	_, err := c.Do("PING")
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Conn() (redisgo.Conn, error) {
	err := Ping(r.pool.Get())
	if err != nil {
		return nil, errors.Wrap(err, "Error while pinging connection for publishing message to redis by channel ")
	}
	return r.pool.Get(), nil
}

// Save temporary date in redis
func (r *RedisClient) SaveTempData(key string, value string) error {
	conn, err := r.Conn()
	if err != nil {
		return errors.Wrap(err, "Error while getting connection for saving temporary data in redis")
	}
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return errors.Wrap(err, "Error while saving temporary data in redis")
	}
	return nil
}

//check if key exists in redis
func (r *RedisClient) KeyExists(key string) (bool, error) {
	conn, err := r.Conn()
	if err != nil {
		return false, errors.Wrap(err, "Error while getting connection for checking key exists in redis")
	}
	defer conn.Close()
	exists, err := redisgo.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, errors.Wrap(err, "Error while checking key exists in redis")
	}
	return exists, nil
}

// Get temporary data from redis
func (r *RedisClient) GetTempData(key string) (string, error) {
	conn, err := r.Conn()
	if err != nil {
		return "", errors.Wrap(err, "Error while getting connection for getting temporary data from redis")
	}
	defer conn.Close()
	value, err := redisgo.String(conn.Do("GET", key))
	if err != nil {
		return "", errors.Wrap(err, "Error while getting temporary data from redis")
	}
	return value, nil
}

//delete temporary data from redis
func (r *RedisClient) DeleteTempData(key string) error {
	conn, err := r.Conn()
	if err != nil {
		return errors.Wrap(err, "Error while getting connection for deleting temporary data from redis")
	}
	defer conn.Close()
	_, err = conn.Do("DEL", key)
	if err != nil {
		return errors.Wrap(err, "Error while deleting temporary data from redis")
	}
	return nil
}