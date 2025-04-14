package routes

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/cmd/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	clientes := r.Group("/clientes")
	{
		clientes.GET("", controller.GetAllClientes)
		clientes.GET("/:id", controller.GetClienteByID)
		clientes.POST("", controller.CreateCliente)
		clientes.PUT("/:id", controller.UpdateCliente)
		clientes.DELETE("/:id", controller.DeleteCliente)
	}

	enderecos := r.Group("/enderecos")
	{
		enderecos.GET("", controller.GetAllEnderecos)
		enderecos.GET("/:id", controller.GetEnderecoByID)
		enderecos.POST("", controller.CreateEndereco)
		enderecos.PUT("/:id", controller.UpdateEndereco)
		enderecos.DELETE("/:id", controller.DeleteEndereco)
	}

	agendamentos := r.Group("/agendamentos")
	{
		agendamentos.GET("", controller.GetAllAgendamentos)
		agendamentos.GET("/:id", controller.GetAgendamentoByID)
		agendamentos.POST("", controller.CreateAgendamento)
		agendamentos.PUT("/:id", controller.UpdateAgendamento)
		agendamentos.DELETE("/:id", controller.DeleteAgendamento)
	}

	orcamentos := r.Group("/orcamentos")
	{
		orcamentos.GET("", controller.GetAllOrcamentos)
		orcamentos.GET("/:id", controller.GetOrcamentoByID)
		orcamentos.POST("", controller.CreateOrcamento)
		orcamentos.PUT("/:id", controller.UpdateOrcamento)
		orcamentos.DELETE("/:id", controller.DeleteOrcamento)
	}

	admin := r.Group("/admin")
	{
		admin.GET("", controller.GetAllAdmins)
		admin.GET("/:id", controller.GetAdminByID)
		admin.POST("", controller.CreateAdmin)
		admin.PUT("/:id", controller.UpdateAdmin)
		admin.DELETE("/:id", controller.DeleteAdmin)
	}
}
