package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/xenakil/cheddartrack/internal/config"
	"github.com/xenakil/cheddartrack/internal/model"
)

var DB *gorm.DB

func Init(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected!")

	err = DB.AutoMigrate(&model.User{}, &model.Category{}, &model.Transaction{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	log.Println("Database migrated!")
}
