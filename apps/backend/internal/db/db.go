package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"panelium/backend/internal/config"
	"panelium/backend/internal/model"
	"sync"
	"time"
)

var (
	initOnce sync.Once
	db       *gorm.DB
)

// TODO: dont allow data creation without ID

func Init() error {
	var err error
	initOnce.Do(func() {
		if db, err = gorm.Open(sqlite.Open(config.DatabaseLocation), &gorm.Config{
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

		if err = db.AutoMigrate(
			&model.Blueprint{},
			&model.Location{},
			&model.Node{},
			&model.NodeAllocation{},
			&model.Server{},
			&model.ServerUser{},
			&model.User{},
			&model.UserMFA{},
			&model.UserMFASession{},
			&model.UserPasswordResetSession{},
			&model.UserSession{},
		); err != nil {
			return
		}
	})

	return err
}

func Instance() *gorm.DB {
	if db == nil {
		panic("database not initialized, call Init() first")
	}
	return db
}

func DeleteExpiredSessions() error {
	db := Instance()
	return db.Where("expiration < ?", time.Now()).Delete(&model.UserSession{}).Error
}
