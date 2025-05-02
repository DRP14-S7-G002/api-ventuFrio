package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {

	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado ou não pôde ser carregado")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbName)

	var db *gorm.DB
	var err error

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Tentando conectar ao banco de dados... (%d/%d)", i+1, maxRetries)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("falha ao conectar com o banco: %w", err)
	}

	DB = db

	// if err != nil {
	// 	log.Printf("Erro ao realizar AutoMigrate: %v", err)
	// 	return err
	// }

	log.Println("Banco conectado e migrado com sucesso.")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
