package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func MustNewMysql(dsn string, maxOpen int, maxIdle int, logDebug bool) (db *gorm.DB) {
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxOpenConns(maxOpen)
	db.DB().SetMaxIdleConns(maxIdle)
	db.LogMode(logDebug)

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	return
}

func NewMysql(dsn string, maxOpen int, maxIdle int, logDebug bool) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}

	db.DB().SetMaxOpenConns(maxOpen)
	db.DB().SetMaxIdleConns(maxIdle)
	db.LogMode(logDebug)
	if err = db.DB().Ping(); err != nil {
		return
	}
	return
}
