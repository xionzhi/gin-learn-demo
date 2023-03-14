package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBSession *gorm.DB
)

func InitPostgresDB(dbConfig string) (*gorm.DB, error) {
	// conn postgres
	var err error
	DBSession, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	return DBSession, err
}

func InitMysqlDB(dbConfig string) (*gorm.DB, error) {
	// conn mysql
	var err error
	DBSession, err = gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	return DBSession, err
}

func InitSqliteDB(dbConfig string) (*gorm.DB, error) {
	// conn sqlite
	var err error
	DBSession, err = gorm.Open(sqlite.Open(dbConfig), &gorm.Config{})
	return DBSession, err
}
