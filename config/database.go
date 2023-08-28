package config

import (
	"fmt"
	"gorm.io/gorm/logger"
	"notes-api/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.ErrorPanic(err)

	fmt.Println("connected to database successfully!")
	return db
}
