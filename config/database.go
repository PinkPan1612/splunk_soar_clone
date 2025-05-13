package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
)

func getEnv() (string, error) {
	required := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSLMODE",
	}

	for _, env := range required {
		if os.Getenv(env) == "" {
			return "", fmt.Errorf("missing required environment variable: %s", env)
		}
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	return connStr, nil
}

func ConnectionDatabase() *gorm.DB {
	once.Do(func() {
		dsn, err := getEnv()
		if err != nil {
			log.Fatalf("❌ Environment error: %v", err)
		}

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatalf("❌ Failed to connect to database: %v", err)
		}
		log.Println("✅ Connected to database")
	})
	return db
}
