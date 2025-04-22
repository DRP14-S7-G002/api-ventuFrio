package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type OrcamentoHandler struct {
	service service.OrcamentoService
}

func NewOrcamentoHandler(s service.OrcamentoService) *OrcamentoHandler {
	return &OrcamentoHandler{s}
}

// GetAllOrcamentos godoc
// @Summary Lista todos os orçamentos
// @Description Retorna todos os registros de orçamentos cadastrados
// @Tags orcamentos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Orcamento
// @Failure 500 {object} map[string]string
// @Router /orcamentos [get]
func (h *OrcamentoHandler) GetAllOrcamentos(c *gin.Context) {
	orcamentos, err := h.service.GetAllOrcamentos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting orcamentos"})
		return
	}
	c.JSON(http.StatusOK, orcamentos)
}

// GetOrcamentoByID godoc
// @Summary Busca um orçamento pelo ID
// @Description Retorna os dados de um orçamento com base no ID fornecido
// @Tags orcamentos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Orçamento"
// @Success 200 {object} models.Orcamento
// @Failure 400,404 {object} map[string]string
// @Router /orcamentos/{id} [get]
func (h *OrcamentoHandler) GetOrcamentoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	orcamento, err := h.service.GetOrcamentoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orcamento not found"})
		return
	}
	c.JSON(http.StatusOK, orcamento)
}

// CreateOrcamento godoc
// @Summary Cria um novo orçamento
// @Description Registra um novo orçamento com os dados fornecidos
// @Tags orcamentos
// @Accept  json
// @Produce  json
// @Param orcamento body models.Orcamento true "Dados do Orçamento"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /orcamentos [post]
func (h *OrcamentoHandler) CreateOrcamento(c *gin.Context) {
	var orcamento models.Orcamento
	if err := c.ShouldBindJSON(&orcamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.CreateOrcamento(orcamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create orcamento"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Orcamento created successfully"})
}

// UpdateOrcamento godoc
// @Summary Atualiza um orçamento existente
// @Description Atualiza os dados de um orçamento com base no ID
// @Tags orcamentos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Orçamento"
// @Param orcamento body models.Orcamento true "Novos dados do Orçamento"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /orcamentos/{id} [put]
func (h *OrcamentoHandler) UpdateOrcamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var orcamento models.Orcamento
	if err := c.ShouldBindJSON(&orcamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.UpdateOrcamento(id, orcamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Orcamento updated"})
}

// DeleteOrcamento godoc
// @Summary Exclui um orçamento
// @Description Remove um orçamento com base no ID fornecido
// @Tags orcamentos
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Orçamento"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /orcamentos/{id} [delete]
func (h *OrcamentoHandler) DeleteOrcamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteOrcamento(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Orcamento deleted"})
}
