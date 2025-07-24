package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func NewDBConnection(config DBConfig) (*gorm.DB, error) {
	dsn := buildDSN(config)
	fmt.Print(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func buildDSN(config DBConfig) string {
	return "host=" + config.Host +
		" user=" + config.User +
		" password=" + config.Password +
		" dbname=" + config.DBName +
		" port=" + config.Port +
		" sslmode=" + config.SSLMode +
		" TimeZone=" + config.TimeZone
}
