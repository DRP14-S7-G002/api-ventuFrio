package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/dto"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type AgendamentoHandler struct {
	service service.AgendamentoService
}

func NewAgendamentoHandler(s service.AgendamentoService) *AgendamentoHandler {
	return &AgendamentoHandler{service: s}
}

// GetAllAgendamentos godoc
// @Summary Lista todos os agendamentos
// @Description Retorna todos os registros de agendamento
// @Tags agendamentos
// @Accept json
// @Produce json
// @Success 200 {array} dto.AgendamentoResponse
// @Failure 500 {object} map[string]string
// @Router /agendamentos [get]
func (h *AgendamentoHandler) GetAllAgendamentos(c *gin.Context) {
	agendamentos, err := h.service.GetAllAgendamentos()
	if err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar agendamentos"})
		return
	}

	var response []dto.AgendamentoResponse
	for _, agendamento := range agendamentos {
		response = append(response, dto.AgendamentoResponse{
			ID:         agendamento.ID,
			DataVisita: agendamento.DataVisita,
			ClienteID:  agendamento.ClienteID,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetAgendamentoByID godoc
// @Summary Busca um agendamento pelo ID
// @Description Retorna um agendamento com base no ID informado
// @Tags agendamentos
// @Accept json
// @Produce json
// @Param id path int true "ID do Agendamento"
// @Success 200 {object} dto.AgendamentoResponse
// @Failure 400,404 {object} map[string]string
// @Router /agendamentos/{id} [get]
func (h *AgendamentoHandler) GetAgendamentoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	agendamento, err := h.service.GetAgendamentoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	response := dto.AgendamentoResponse{
		ID:         agendamento.ID,
		DataVisita: agendamento.DataVisita,
		ClienteID:  agendamento.ClienteID,
	}

	c.JSON(http.StatusOK, response)
}

// CreateAgendamento godoc
// @Summary Cria um novo agendamento
// @Description Cria um novo agendamento com os dados fornecidos
// @Tags agendamentos
// @Accept json
// @Produce json
// @Param agendamento body dto.AgendamentoCreateRequest true "Dados do Agendamento"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /agendamentos [post]
func (h *AgendamentoHandler) CreateAgendamento(c *gin.Context) {
	var req dto.AgendamentoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	agendamento := models.Agendamento{
		DataVisita: req.DataVisita,
		ClienteID:  req.ClienteID,
	}

	if err := h.service.CreateAgendamento(agendamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível criar o agendamento"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Agendamento criado com sucesso"})
}

// UpdateAgendamento godoc
// @Summary Atualiza um agendamento existente
// @Description Atualiza os dados de um agendamento pelo ID
// @Tags agendamentos
// @Accept json
// @Produce json
// @Param id path int true "ID do Agendamento"
// @Param agendamento body dto.AgendamentoCreateRequest true "Novos dados do Agendamento"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /agendamentos/{id} [put]
func (h *AgendamentoHandler) UpdateAgendamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.AgendamentoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	agendamento := models.Agendamento{
		DataVisita: req.DataVisita,
		ClienteID:  req.ClienteID,
	}

	if err := h.service.UpdateAgendamento(id, agendamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar agendamento"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Agendamento atualizado com sucesso"})
}

// DeleteAgendamento godoc
// @Summary Remove um agendamento
// @Description Exclui um agendamento com base no ID fornecido
// @Tags agendamentos
// @Accept json
// @Produce json
// @Param id path int true "ID do Agendamento"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /agendamentos/{id} [delete]
func (h *AgendamentoHandler) DeleteAgendamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteAgendamento(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao excluir agendamento"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Agendamento excluído com sucesso"})
}
