package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	strConn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, name, port,
	)
	conn, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	db = conn
	config, _ := db.DB()

	config.SetConnMaxLifetime(10)
	config.SetMaxOpenConns(100)
	config.SetMaxIdleConns(10)

	// migrations.RunMigrations(db)
	log.Println("✅ Banco de dados conectado com sucesso.")
}

func GetDatabase() *gorm.DB {
	return db
}
