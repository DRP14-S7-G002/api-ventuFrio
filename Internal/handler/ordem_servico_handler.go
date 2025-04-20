package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type OrdemDeServicoHandler struct {
	service service.OrdemDeServicoService
}

func NewOrdemDeServicoHandler(s service.OrdemDeServicoService) *OrdemDeServicoHandler {
	return &OrdemDeServicoHandler{s}
}

// GetAllOrdens godoc
// @Summary Lista todas as ordens de serviço
// @Description Retorna todas as ordens de serviço cadastradas
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Success 200 {array} models.OrdemDeServico
// @Failure 500 {object} map[string]string
// @Router /ordens [get]
func (h *OrdemDeServicoHandler) GetAllOrdens(c *gin.Context) {
	ordens, err := h.service.GetAllOrdens()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving work ordens"})
		return
	}
	c.JSON(http.StatusOK, ordens)
}

// GetOrdemByID godoc
// @Summary Busca uma ordem de serviço pelo ID
// @Description Retorna os dados da ordem de serviço com base no ID fornecido
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Param id path int true "ID da Ordem de Serviço"
// @Success 200 {object} models.OrdemDeServico
// @Failure 400,404 {object} map[string]string
// @Router /ordens/{id} [get]
func (h *OrdemDeServicoHandler) GetOrdemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	ordem, err := h.service.GetOrdemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Work ordem not found"})
		return
	}
	c.JSON(http.StatusOK, ordem)
}

// CreateOrdem godoc
// @Summary Cria uma nova ordem de serviço
// @Description Registra uma nova ordem de serviço com os dados fornecidos
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Param ordem body models.OrdemDeServico true "Dados da Ordem de Serviço"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /ordens [post]
func (h *OrdemDeServicoHandler) CreateOrdem(c *gin.Context) {
	var ordem models.OrdemDeServico
	if err := c.ShouldBindJSON(&ordem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.CreateOrdem(ordem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create work ordem"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Work ordem created successfully"})
}

// UpdateOrdem godoc
// @Summary Atualiza uma ordem de serviço existente
// @Description Atualiza os dados de uma ordem de serviço com base no ID
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Param id path int true "ID da Ordem de Serviço"
// @Param ordem body models.OrdemDeServico true "Novos dados da Ordem de Serviço"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /ordens/{id} [put]
func (h *OrdemDeServicoHandler) UpdateOrdem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var ordem models.OrdemDeServico
	if err := c.ShouldBindJSON(&ordem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.UpdateOrdem(id, ordem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Work ordem updated"})
}

// DeleteOrdem godoc
// @Summary Exclui uma ordem de serviço
// @Description Remove uma ordem de serviço com base no ID fornecido
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Param id path int true "ID da Ordem de Serviço"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /ordens/{id} [delete]
func (h *OrdemDeServicoHandler) DeleteOrdem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteOrdem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Work ordem deleted"})
}
