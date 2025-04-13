package db

import (
	"fmt"
	"log"
	"os"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	err = DB.AutoMigrate(
		&models.Cliente{},
		&models.Endereco{},
		&models.Agendamento{},
		&models.Orcamento{},
		&models.Admin{},
	)
	if err != nil {
		log.Printf("Erro ao realizar AutoMigrate: %v", err)
		return err
	}

	return nil
}
