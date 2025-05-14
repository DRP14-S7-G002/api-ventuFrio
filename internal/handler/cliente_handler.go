package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/dto"
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
// @Accept json
// @Produce json
// @Success 200 {array} dto.ClienteResponse
// @Failure 500 {object} map[string]string
// @Router /clientes [get]
func (h *ClienteHandler) GetAllClientes(c *gin.Context) {
	clientes, err := h.service.GetAllClientes()
	if err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
		return
	}

	var response []dto.ClienteResponse
	for _, cliente := range clientes {
		response = append(response, dto.ClienteResponse{
			ID:       cliente.ID,
			Nome:     cliente.Nome,
			Telefone: cliente.Telefone,
			CPF:      cliente.CPF,
			Rua:      cliente.Rua,
			Numero:   cliente.Numero,
			Bairro:   cliente.Bairro,
			CEP:      cliente.CEP,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetClienteByID godoc
// @Summary Busca um cliente pelo ID
// @Description Retorna um cliente com base no ID informado
// @Tags clientes
// @Accept json
// @Produce json
// @Param id path int true "ID do Cliente"
// @Success 200 {object} dto.ClienteResponse
// @Failure 400,404 {object} map[string]string
// @Router /clientes/{id} [get]
func (h *ClienteHandler) GetClienteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	cliente, err := h.service.GetClienteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	response := dto.ClienteResponse{
		ID:       cliente.ID,
		Nome:     cliente.Nome,
		Telefone: cliente.Telefone,
		CPF:      cliente.CPF,
		Rua:      cliente.Rua,
		Numero:   cliente.Numero,
		Bairro:   cliente.Bairro,
		CEP:      cliente.CEP,
	}

	c.JSON(http.StatusOK, response)
}

// CreateCliente godoc
// @Summary Cria um novo cliente
// @Description Cria um novo cliente com os dados fornecidos
// @Tags clientes
// @Accept json
// @Produce json
// @Param cliente body dto.ClienteCreateRequest true "Dados do Cliente"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /clientes [post]
func (h *ClienteHandler) CreateCliente(c *gin.Context) {
	var req dto.ClienteCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Requisição inválida"})
		return
	}

	cliente := models.Cliente{
		Nome:     req.Nome,
		Telefone: req.Telefone,
		CPF:      req.CPF,
		Rua:      req.Rua,
		Numero:   req.Numero,
		Bairro:   req.Bairro,
		CEP:      req.CEP,
	}

	if err := h.service.CreateCliente(cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar cliente"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Cliente criado com sucesso"})
}

// UpdateCliente godoc
// @Summary Atualiza um cliente existente
// @Description Atualiza os dados de um cliente pelo ID
// @Tags clientes
// @Accept json
// @Produce json
// @Param id path int true "ID do Cliente"
// @Param cliente body dto.ClienteCreateRequest true "Novos dados do Cliente"
// @Success 200 {object} map[string]string
// @Failure 400,404,500 {object} map[string]string
// @Router /clientes/{id} [put]
func (h *ClienteHandler) UpdateCliente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.ClienteCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	cliente := models.Cliente{
		Nome:     req.Nome,
		Telefone: req.Telefone,
		CPF:      req.CPF,
		Rua:      req.Rua,
		Numero:   req.Numero,
		Bairro:   req.Bairro,
		CEP:      req.CEP,
	}

	if err := h.service.UpdateCliente(id, cliente); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar cliente"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cliente atualizado com sucesso"})
}

// DeleteCliente godoc
// @Summary Remove um cliente
// @Description Exclui um cliente com base no ID fornecido
// @Tags clientes
// @Accept json
// @Produce json
// @Param id path int true "ID do Cliente"
// @Success 200 {object} map[string]string
// @Failure 400,404,500 {object} map[string]string
// @Router /clientes/{id} [delete]
func (h *ClienteHandler) DeleteCliente(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteCliente(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao excluir cliente"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cliente excluído com sucesso"})
}
