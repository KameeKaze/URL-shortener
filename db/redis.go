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
			Addr:     "storage:6379",
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		}),
	}
)

func (r *REDIS) SetURL(URI, URL string) error {
	err := r.db.Set(URI, URL, 0).Err()
	return err
}

func (r *REDIS) GetURL(URL string) (string, error) {
	val, err := r.db.Get(URL).Result()
	return val, err
}
