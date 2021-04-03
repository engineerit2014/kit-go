package pgdb

import (
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var pgOptions = &pg.Options{
	Addr:     "db-test:5432",
	User:     "root",
	Password: "root",
	Database: "testdb",
}

func TestNewPgDB_WhenOptionsOk_ThenCreateClient(t *testing.T) {
	got := NewPgDB(pgOptions)
	defer got.Close()

	assert.NotNil(t, got)
	assert.Equal(t, pgOptions.Addr, got.Options().Addr)
	assert.Equal(t, pgOptions.User, got.Options().User)
	assert.Equal(t, pgOptions.Password, got.Options().Password)
	assert.Equal(t, pgOptions.Database, got.Options().Database)
}
