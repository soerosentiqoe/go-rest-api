package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDb is for koneksi
func SetupDb() *gorm.DB {
	errenv := godotenv.Load()
	if errenv != nil {
		panic("error")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3314)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
	}
	return db
}

//CloseDb untuk close koneksi
func CloseDb(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("error close coneksi")
	}
	dbSQL.Close()
}
