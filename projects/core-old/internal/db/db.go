package db

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	initOnce sync.Once
	db       *gorm.DB
)

func Instance() (*gorm.DB, error) {
	var err error

	initOnce.Do(func() { // TODO: make db location configurable
		if db, err = gorm.Open(sqlite.Open("./db.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}); err != nil {
			return
		}

		if tx := db.Exec("PRAGMA journal_mode=WAL;"); tx.Error != nil {
			err = tx.Error
			return
		}
		if tx := db.Exec("PRAGMA synchronous=NORMAL;"); tx.Error != nil {
			err = tx.Error
			return
		}

		if err = db.AutoMigrate(); err != nil {
			return
		}
	})

	if db == nil {
		panic("database failed to initialize")
	}

	return db, err
}
