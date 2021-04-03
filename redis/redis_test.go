package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

var redisOptions = &redis.Options{
	Addr:     "redis-test:6379",
	Password: "",
	DB:       0,
}

func TestNewRedisDB_WhenOptionsOk_ThenCreateClient(t *testing.T) {
	got := NewRedisDB(redisOptions)
	got.Close()

	assert.NotNil(t, got)
	assert.Equal(t, redisOptions.Addr, got.Options().Addr)
	assert.Equal(t, redisOptions.Password, got.Options().Password)
	assert.Equal(t, redisOptions.DB, got.Options().DB)
}
