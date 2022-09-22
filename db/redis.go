package db

import (
	"os"

	"github.com/go-redis/redis"
)

type REDIS struct {
	db *redis.Client
}

// define database connection for redis
var (
	Redis = REDIS{
		db: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		}),
	}
)
