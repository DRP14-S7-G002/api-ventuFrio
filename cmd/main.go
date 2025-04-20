// @title API VentuFrio
// @version 1.0
// @description API para gerenciamento de clientes, agendamentos, orçamentos e ordens de serviço.
// @termsOfService http://swagger.io/terms/

// @contact.name Suporte VentuFrio
// @contact.email suporte@ventufrio.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"log"
	"os"

	"github.com/DRP14-S7-G002/api-ventuFrio/cmd/routes"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"github.com/DRP14-S7-G002/api-ventuFrio/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	docs.SwaggerInfo.BasePath = "/api/v1"

	err := db.InitDB()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
		os.Exit(1)
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
