package main

import (
	"context"
	"github.com/Lairon/db-go/redisdb"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"time"
)

type RedisClient struct {
	db *redis.Client
}

func main() {
	client := NewRedisClient()
	defer client.db.Close()

	ctx := context.Background()

	key := "key1"
	value := "val1"

	client.SetKey(ctx, key, value, 10*time.Second)
	val := client.GetKey(ctx, key)

	log.Infof("Key: %s, Value: %s", key, val)
}

func NewRedisClient() *RedisClient {
	return &RedisClient{
		db: redisdb.NewRedisDB(&redis.Options{
			Addr:     "redis-test:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func (r *RedisClient) SetKey(ctx context.Context, key string, val string, expiration time.Duration) {
	if err := r.db.Set(ctx, key, val, expiration).Err(); err != nil {
		log.Panicf("Redis error [%+v] in set func", err)
	}
}

func (r *RedisClient) GetKey(ctx context.Context, key string) string {
	val, err := r.db.Get(ctx, key).Result()
	if err != nil {
		log.Panicf("Redis error [%+v] in get func", err)
	}

	return val
}
