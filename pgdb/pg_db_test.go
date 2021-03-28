package pgdb

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var pgOptions = &pg.Options{
	Addr:     "localhost:5432",
	User:     "postgres",
	Password: "postgres",
	Database: "pg-db-go",
}

func TestNewPgDB_WhenOptionsOk_ThenCreateClient(t *testing.T) {
	got := NewPgDB(pgOptions)
	got.Close()

	assert.NotNil(t, got)
	assert.Equal(t, pgOptions.Addr, got.Options().Addr)
	assert.Equal(t, pgOptions.User, got.Options().User)
	assert.Equal(t, pgOptions.Password, got.Options().Password)
	assert.Equal(t, pgOptions.Database, got.Options().Database)
}

func TestNewPgDB_WhenOptionsNil_ThenPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, fmt.Sprint("runtime error: invalid memory address or nil pointer dereference"), fmt.Sprint(r))
		}
	}()
	NewPgDB(nil)
}

func TestNewPgDB_WhenOptionsContainBadAddr_ThenPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, fmt.Sprintln("Postgres connection error dial tcp: lookup bad: no such host"), fmt.Sprint(r))
		}
	}()
	NewPgDB(&pg.Options{
		Addr:     "bad:6379",
		User:     "postgres",
		Password: "postgres",
		Database: "pg-db-go",
	})
}
