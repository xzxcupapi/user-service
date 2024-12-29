package database

import (
	"fmt"
	"user-services/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	config := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}

	db, err := gorm.Open(postgres.New(config), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}
