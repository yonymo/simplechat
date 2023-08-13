package db

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/yonymo/simplechat/pkg/options"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetDBInstance(opts *options.MySQLOptions) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		if opts == nil {
			err = errors.New("mysql options is nil")
			return
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", opts.UserName, opts.Password, opts.Host, opts.Port, opts.Database)
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.LogLevel(opts.LogLevel),
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})

		var sqlDB *sql.DB
		sqlDB, err = db.DB()

		sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
		sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
		sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifetime)

		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		return nil, errors.New("mysql db open failed")
	}
	if db == nil {
		panic("db is nil")
	}
	return db, nil

}
