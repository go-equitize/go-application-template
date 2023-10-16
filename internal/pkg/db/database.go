package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	cfg "go-template/internal/pkg/config"
)

var db *gorm.DB

func InitDB() error {
	if db != nil {
		return nil
	}

	dbCfg := cfg.Instance().DB
	gormDB, err := gorm.Open(
		mysql.Open(dbCfg.DSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.LogLevel(dbCfg.LogLevel)),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to open connection, err: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("failed to get *sql.db, err: %v", err)
	}

	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.ConnLifeTime) * time.Second)

	if err = sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database, err: %v", err)
	}

	fmt.Println("Successfully connect to Mysql database.")

	db = gormDB
	return nil
}

func Instance() *gorm.DB {
	return db
}
