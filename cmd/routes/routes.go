package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/handler"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/repository"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
)

func RegisterRoutes(r *gin.Engine) {
	dbInstance := db.GetDB()

	clienteRepo := repository.NewClienteRepository(dbInstance)
	clienteService := service.NewClienteService(clienteRepo)
	clienteHandler := handler.NewClienteHandler(clienteService)

	agendamentoRepo := repository.NewAgendamentoRepository(dbInstance)
	agendamentoService := service.NewAgendamentoService(agendamentoRepo)
	agendamentoHandler := handler.NewAgendamentoHandler(agendamentoService)

	orcamentoRepo := repository.NewOrcamentoRepository(dbInstance)
	orcamentoService := service.NewOrcamentoService(orcamentoRepo)
	orcamentoHandler := handler.NewOrcamentoHandler(orcamentoService)

	ordemRepo := repository.NewOrdemDeServicoRepository(dbInstance)
	ordemService := service.NewOrdemDeServicoService(ordemRepo)
	ordemHandler := handler.NewOrdemDeServicoHandler(ordemService)

	materialRepo := repository.NewMaterialRepository(dbInstance)
	materialService := service.NewMaterialService(materialRepo)
	materialHandler := handler.NewMaterialHandler(materialService)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/clientes", clienteHandler.GetAllClientes)
		v1.GET("/clientes/:id", clienteHandler.GetClienteByID)
		v1.POST("/clientes", clienteHandler.CreateCliente)
		v1.PUT("/clientes/:id", clienteHandler.UpdateCliente)
		v1.DELETE("/clientes/:id", clienteHandler.DeleteCliente)

		v1.GET("/agendamentos", agendamentoHandler.GetAllAgendamentos)
		v1.GET("/agendamentos/:id", agendamentoHandler.GetAgendamentoByID)
		v1.POST("/agendamentos", agendamentoHandler.CreateAgendamento)
		v1.PUT("/agendamentos/:id", agendamentoHandler.UpdateAgendamento)
		v1.DELETE("/agendamentos/:id", agendamentoHandler.DeleteAgendamento)

		v1.GET("/orcamentos", orcamentoHandler.GetAllOrcamentos)
		v1.GET("/orcamentos/:id", orcamentoHandler.GetOrcamentoByID)
		v1.POST("/orcamentos", orcamentoHandler.CreateOrcamento)
		v1.PUT("/orcamentos/:id", orcamentoHandler.UpdateOrcamento)
		v1.DELETE("/orcamentos/:id", orcamentoHandler.DeleteOrcamento)

		v1.GET("/ordens", ordemHandler.GetAllOrdens)
		v1.GET("/ordens/:id", ordemHandler.GetOrdemByID)
		v1.POST("/ordens", ordemHandler.CreateOrdem)
		v1.PUT("/ordens/:id", ordemHandler.UpdateOrdem)
		v1.DELETE("/ordens/:id", ordemHandler.DeleteOrdem)

		v1.GET("/materiais", materialHandler.GetAllMateriais)
		v1.GET("/materiais/:id", materialHandler.GetMaterialByID)
		v1.POST("/materiais", materialHandler.CreateMaterial)
		v1.PUT("/materiais/:id", materialHandler.UpdateMaterial)
		v1.DELETE("/materiais/:id", materialHandler.DeleteMaterial)
	}
}
