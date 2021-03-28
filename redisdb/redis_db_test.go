package redisdb

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

var redisOptions = &redis.Options{
	Addr:     "localhost:6379",
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

func TestNewRedisDB_WhenOptionsNil_ThenPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, fmt.Sprint("runtime error: invalid memory address or nil pointer dereference"), fmt.Sprint(r))
		}
	}()
	NewRedisDB(nil)
}

func TestNewRedisDB_WhenOptionsContainBadAddr_ThenPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, fmt.Sprintln("Redis connection error dial tcp: lookup awdawdaw: no such host"), fmt.Sprint(r))
		}
	}()
	NewRedisDB(&redis.Options{
		Addr:     "bad:6379",
		Password: "",
		DB:       0,
	})
}
