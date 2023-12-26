package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {
	errENV := godotenv.Load()
	if errENV != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect database")
	}

	defer DisconnectDB(db)

	return db

}

// DisconnectDB is stopping your connection to postgrsql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}

func ConnectDBGen() *sql.DB {

	errENV := godotenv.Load()
	if errENV != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()

	return db
}
