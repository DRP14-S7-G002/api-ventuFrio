package main

import (
	"log"
	"os"

	"github.com/DRP14-S7-G002/api-ventuFrio/cmd/routes"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	err := db.InitDB()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
		os.Exit(1)
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
