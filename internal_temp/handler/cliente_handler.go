package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	service service.ClienteService
}

func NewClienteHandler(s service.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: s}
}

// GetAllClientes godoc
// @Summary Lista todos os clientes
// @Description Retorna todos os registros de clientes
// @Tags clientes
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Cliente
// @Failure 500 {object} map[string]string
// @Router /clientes [get]
func (h *ClienteHandler) GetAllClientes(c *gin.Context) {
	clientes, err := h.service.GetAllClientes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving clientes"})
		return
	}
	c.JSON(http.StatusOK, clientes)
}

// GetClienteByID godoc
// @Summary Busca um cliente pelo ID
// @Description Retorna um cliente com base no ID informado
// @Tags clientes
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Cliente"
// @Success 200 {object} models.Cliente
// @Failure 400,404 {object} map[string]string
// @Router /clientes/{id} [get]
func (h *ClienteHandler) GetClienteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	cliente, err := h.service.GetClienteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente not found"})
		return
	}
	c.JSON(http.StatusOK, cliente)
}

// CreateCliente godoc
// @Summary Cria um novo cliente
// @Description Cria um novo cliente com os dados fornecidos
// @Tags clientes
// @Accept  json
// @Produce  json
// @Param cliente body models.Cliente true "Dados do Cliente"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /clientes [post]
func (h *ClienteHandler) CreateCliente(c *gin.Context) {
	var cliente models.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := h.service.CreateCliente(cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cliente"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Cliente created successfully"})
}

// UpdateCliente godoc
// @Summary Atualiza um cliente existente
// @Description Atualiza os dados de um cliente pelo ID
// @Tags clientes
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Cliente"
// @Param cliente body models.Cliente true "Novos dados do Cliente"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /clientes/{id} [put]
func (h *ClienteHandler) UpdateCliente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var cliente models.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.UpdateCliente(id, cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cliente updated"})
}

// DeleteCliente godoc
// @Summary Remove um cliente
// @Description Exclui um cliente com base no ID fornecido
// @Tags clientes
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Cliente"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /clientes/{id} [delete]
func (h *ClienteHandler) DeleteCliente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteCliente(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cliente deleted"})
}
