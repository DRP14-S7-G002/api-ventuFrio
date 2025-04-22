package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type AgendamentoHandler struct {
	service service.AgendamentoService
}

func NewAgendamentoHandler(s service.AgendamentoService) *AgendamentoHandler {
	return &AgendamentoHandler{s}
}

// GetAllAgendamentos godoc
// @Summary Lista todos os agendamentos
// @Description Retorna todos os registros de agendamento
// @Tags agendamentos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Agendamento
// @Failure 500 {object} map[string]string
// @Router /agendamentos [get]
func (h *AgendamentoHandler) GetAllAgendamentos(c *gin.Context) {
	agendamentos, err := h.service.GetAllAgendamentos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting agendamentos"})
		return
	}
	c.JSON(http.StatusOK, agendamentos)
}

// GetAgendamentoByID godoc
// @Summary Busca um agendamento pelo ID
// @Description Retorna um agendamento com base no ID informado
// @Tags agendamentos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Agendamento"
// @Success 200 {object} models.Agendamento
// @Failure 400,404 {object} map[string]string
// @Router /agendamentos/{id} [get]
func (h *AgendamentoHandler) GetAgendamentoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	agendamento, err := h.service.GetAgendamentoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento not found"})
		return
	}
	c.JSON(http.StatusOK, agendamento)
}

// CreateAgendamento godoc
// @Summary Cria um novo agendamento
// @Description Cria um novo agendamento com os dados fornecidos
// @Tags agendamentos
// @Accept  json
// @Produce  json
// @Param agendamento body models.Agendamento true "Dados do Agendamento"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /agendamentos [post]
func (h *AgendamentoHandler) CreateAgendamento(c *gin.Context) {
	var agendamento models.Agendamento
	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.CreateAgendamento(agendamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create agendamento"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Agendamento created successfully"})
}

// UpdateAgendamento godoc
// @Summary Atualiza um agendamento existente
// @Description Atualiza os dados de um agendamento pelo ID
// @Tags agendamentos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Agendamento"
// @Param agendamento body models.Agendamento true "Novos dados do Agendamento"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /agendamentos/{id} [put]
func (h *AgendamentoHandler) UpdateAgendamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var agendamento models.Agendamento
	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.UpdateAgendamento(id, agendamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Agendamento updated"})
}

// DeleteAgendamento godoc
// @Summary Remove um agendamento
// @Description Exclui um agendamento com base no ID fornecido
// @Tags agendamentos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Agendamento"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /agendamentos/{id} [delete]
func (h *AgendamentoHandler) DeleteAgendamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteAgendamento(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Agendamento deleted"})
}
