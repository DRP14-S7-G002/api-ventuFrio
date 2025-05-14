package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/dto"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type OrcamentoHandler struct {
	service service.OrcamentoService
}

func NewOrcamentoHandler(s service.OrcamentoService) *OrcamentoHandler {
	return &OrcamentoHandler{service: s}
}

// GetAllOrcamentos godoc
// @Summary Lista todos os orçamentos
// @Description Retorna todos os registros de orçamentos cadastrados
// @Tags orcamentos
// @Accept json
// @Produce json
// @Success 200 {array} dto.OrcamentoResponse
// @Failure 500 {object} map[string]string
// @Router /orcamentos [get]
func (h *OrcamentoHandler) GetAllOrcamentos(c *gin.Context) {
	orcamentos, err := h.service.GetAllOrcamentos()
	if err != nil {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar orçamentos"})
		return
	}

	var response []dto.OrcamentoResponse
	for _, orcamento := range orcamentos {
		response = append(response, dto.OrcamentoResponse{
			ID:               orcamento.ID,
			DescricaoInicial: orcamento.DescricaoInicial,
			DescricaoItem:    orcamento.DescricaoItem,
			Status:           orcamento.Status,
			PrazoEntrega:     orcamento.PrazoEntrega.Format("2006-01-02"),
			Valor:            orcamento.Valor,
			ClienteID:        orcamento.ClienteID,
			AgendamentoID:    orcamento.AgendamentoID,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetOrcamentoByID godoc
// @Summary Busca um orçamento pelo ID
// @Description Retorna os dados de um orçamento com base no ID fornecido
// @Tags orcamentos
// @Accept json
// @Produce json
// @Param id path int true "ID do Orçamento"
// @Success 200 {object} dto.OrcamentoResponse
// @Failure 400,404 {object} map[string]string
// @Router /orcamentos/{id} [get]
func (h *OrcamentoHandler) GetOrcamentoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	orcamento, err := h.service.GetOrcamentoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orcamento não encontrado"})
		return
	}

	response := dto.OrcamentoResponse{
		ID:               orcamento.ID,
		DescricaoInicial: orcamento.DescricaoInicial,
		DescricaoItem:    orcamento.DescricaoItem,
		Status:           orcamento.Status,
		PrazoEntrega:     orcamento.PrazoEntrega.Format("2006-01-02"),
		Valor:            orcamento.Valor,
		ClienteID:        orcamento.ClienteID,
		AgendamentoID:    orcamento.AgendamentoID,
	}

	c.JSON(http.StatusOK, response)
}

// CreateOrcamento godoc
// @Summary Cria um novo orçamento
// @Description Registra um novo orçamento com os dados fornecidos
// @Tags orcamentos
// @Accept json
// @Produce json
// @Param orcamento body dto.OrcamentoCreateRequest true "Dados do Orçamento"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /orcamentos [post]
func (h *OrcamentoHandler) CreateOrcamento(c *gin.Context) {
	var req dto.OrcamentoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	prazoEntregaParsed, err := time.Parse("2006-01-02", req.PrazoEntrega)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prazo_entrega inválido, deve ser no formato YYYY-MM-DD"})
		return
	}

	orcamento := models.Orcamento{
		DescricaoInicial: req.DescricaoInicial,
		DescricaoItem:    req.DescricaoItem,
		Status:           req.Status,
		PrazoEntrega:     prazoEntregaParsed,
		Valor:            req.Valor,
		ClienteID:        req.ClienteID,
		AgendamentoID:    req.AgendamentoID,
	}

	if err := h.service.CreateOrcamento(orcamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar orçamento"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Orcamento criado com sucesso"})
}

// UpdateOrcamento godoc
// @Summary Atualiza um orçamento existente
// @Description Atualiza os dados de um orçamento com base no ID
// @Tags orcamentos
// @Accept json
// @Produce json
// @Param id path int true "ID do Orçamento"
// @Param orcamento body dto.OrcamentoCreateRequest true "Novos dados do Orçamento"
// @Success 200 {object} map[string]string
// @Failure 400,404,500 {object} map[string]string
// @Router /orcamentos/{id} [put]
func (h *OrcamentoHandler) UpdateOrcamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.OrcamentoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	prazoEntregaParsed, err := time.Parse("2006-01-02", req.PrazoEntrega)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prazo_entrega inválido, deve ser no formato YYYY-MM-DD"})
		return
	}

	orcamento := models.Orcamento{
		DescricaoInicial: req.DescricaoInicial,
		DescricaoItem:    req.DescricaoItem,
		Status:           req.Status,
		PrazoEntrega:     prazoEntregaParsed,
		Valor:            req.Valor,
		ClienteID:        req.ClienteID,
		AgendamentoID:    req.AgendamentoID,
	}

	if err := h.service.UpdateOrcamento(id, orcamento); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar orçamento"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Orcamento atualizado com sucesso"})
}

// DeleteOrcamento godoc
// @Summary Exclui um orçamento
// @Description Remove um orçamento com base no ID fornecido
// @Tags orcamentos
// @Accept json
// @Produce json
// @Param id path int true "ID do Orçamento"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /orcamentos/{id} [delete]
func (h *OrcamentoHandler) DeleteOrcamento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteOrcamento(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao excluir orçamento"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Orcamento excluído com sucesso"})
}
