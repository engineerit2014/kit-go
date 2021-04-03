package redisdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
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
	log.Infoln("Connecting to redis database...")
	ctx := context.Background()
	db := redis.NewClient(redisOptions)

	if err := db.Ping(ctx).Err(); err != nil {
		log.Errorln(err)
		time.Sleep(2 * time.Second)

		db := redis.NewClient(redisOptions)
		if err := db.Ping(ctx).Err(); err != nil {
			log.Panicf("Redis connection error %+v\n", err)
		}
	}

	log.Infoln("Redis successfully connected...")
	return db
}
