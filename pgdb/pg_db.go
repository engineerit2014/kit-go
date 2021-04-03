package pgdb

import (
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
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
	log.Infoln("Connecting to postgres database...")
	db := pg.Connect(pgOptions)

	var n int
	if _, err := db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
		log.Errorln(err)
		time.Sleep(2 * time.Second)

		db = pg.Connect(pgOptions)
		if _, err := db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
			log.Panicf("Postgres connection error %+v\n", err)
		}
	}

	log.Infoln("Connection to postgres verified and successfully connected...")
	return db
}
