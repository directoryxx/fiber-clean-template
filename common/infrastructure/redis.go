package infrastructure

import (
	"os"

	"github.com/go-redis/redis/v8"
)

var clientRedis *redis.Client

// RedisInit - Initialize the redis client
func RedisInit() *redis.Client {
	dsn := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	clientRedis = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	return clientRedis
}
