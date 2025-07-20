package database

import (
	"fmt"
	"log"
	"role-based/config/env"
	"role-based/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DbConnection() error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.Config("DB_HOST"), env.Config("DB_USER"), env.Config("DB_PASSWORD"), env.Config("DB_NAME"), env.Config("DB_PORT"))

	dabatabse, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,                               // True for productiohn
		Logger:                 logger.Default.LogMode(logger.Info), //LogMode(logger.Warn) for Production
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = dabatabse

	//AutoMigrate DB
	//Account Struct
	err = DB.AutoMigrate(&models.Account{})
	if err != nil {
		log.Println("Cannot Create Account Table!")
		return err
	}

	//Task Struct
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Println("Cannot Create Task Table!")
		return err
	}

	log.Println("Database connected successfully!")
	log.Println("Database migration complete.")
	return nil
}
