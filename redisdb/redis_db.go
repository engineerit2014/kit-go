package redisdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

var (
	instance *redis.Client
	once     sync.Once
)

func NewRedisDB(redisOptions *redis.Options) *redis.Client {
	once.Do(func() {
		instance = Connect(redisOptions)
	})

	return instance
}

func Connect(redisOptions *redis.Options) *redis.Client {
	db := redis.NewClient(redisOptions)

	if err := db.Ping(context.Background()).Err(); err != nil {
		log.Panicf("Redis connection error %+v\n", err)
	}

	log.Println("Redis successfully connected to ->", db.Options().Addr)
	return db
}
