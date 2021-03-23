package pgdb

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"sync"
)

var (
	instance *pg.DB
	once     sync.Once
)

func NewPgDB(pgOptions *pg.Options) *pg.DB {
	once.Do(func() {
		instance = Connect(pgOptions)
	})

	return instance
}

func Connect(pgOptions *pg.Options) *pg.DB {
	db := pg.Connect(pgOptions)

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		fmt.Errorf("Connection error %+v\n", err)
		panic(err)
	}

	fmt.Println("Successfully connected to ->", db.Options().Addr)

	return db
}