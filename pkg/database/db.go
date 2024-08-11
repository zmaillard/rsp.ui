package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"highway-sign-portal-builder/pkg/config"
)

func InitializeDatabase(cfg config.Config) (*gorm.DB, error) {
	var dsn string
	if cfg.DBPassword == "" {
		dsn = fmt.Sprintf("host=%s user=%s  dbname=%s port=%s  search_path=sign", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort)

	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require search_path=sign", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
