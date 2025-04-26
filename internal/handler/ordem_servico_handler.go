package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/dto"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type OrdemDeServicoHandler struct {
	service service.OrdemDeServicoService
}

func NewOrdemDeServicoHandler(s service.OrdemDeServicoService) *OrdemDeServicoHandler {
	return &OrdemDeServicoHandler{service: s}
}

// GetAllOrdens godoc
// @Summary Lista todas as ordens de serviço
// @Description Retorna todas as ordens de serviço cadastradas
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Success 200 {array} dto.OrdemServicoResponse
// @Failure 500 {object} map[string]string
// @Router /ordens [get]
func (h *OrdemDeServicoHandler) GetAllOrdens(c *gin.Context) {
	ordens, err := h.service.GetAllOrdens()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar ordens de serviço"})
		return
	}

	var response []dto.OrdemServicoResponse
	for _, ordem := range ordens {
		response = append(response, dto.OrdemServicoResponse{
			ID:               ordem.ID,
			DescricaoServico: ordem.DescricaoServico,
			Status:           ordem.Status,
			Responsavel:      ordem.Responsavel,
			OrcamentoID:      ordem.OrcamentoID,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetOrdemByID godoc
// @Summary Busca uma ordem de serviço pelo ID
// @Description Retorna os dados da ordem de serviço com base no ID fornecido
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Param id path int true "ID da Ordem de Serviço"
// @Success 200 {object} dto.OrdemServicoResponse
// @Failure 400,404 {object} map[string]string
// @Router /ordens/{id} [get]
func (h *OrdemDeServicoHandler) GetOrdemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ordem, err := h.service.GetOrdemByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ordem de serviço não encontrada"})
		return
	}

	response := dto.OrdemServicoResponse{
		ID:               ordem.ID,
		DescricaoServico: ordem.DescricaoServico,
		Status:           ordem.Status,
		Responsavel:      ordem.Responsavel,
		OrcamentoID:      ordem.OrcamentoID,
	}

	c.JSON(http.StatusOK, response)
}

// CreateOrdem godoc
// @Summary Cria uma nova ordem de serviço
// @Description Registra uma nova ordem de serviço com os dados fornecidos
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Param ordem body dto.OrdemServicoCreateRequest true "Dados da Ordem de Serviço"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /ordens [post]
func (h *OrdemDeServicoHandler) CreateOrdem(c *gin.Context) {
	var req dto.OrdemServicoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	ordem := models.OrdemDeServico{
		DescricaoServico: req.DescricaoServico,
		Status:           req.Status,
		Responsavel:      req.Responsavel,
		OrcamentoID:      req.OrcamentoID,
	}

	if err := h.service.CreateOrdem(ordem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar ordem de serviço"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Ordem de serviço criada com sucesso"})
}

// UpdateOrdem godoc
// @Summary Atualiza uma ordem de serviço existente
// @Description Atualiza os dados de uma ordem de serviço com base no ID
// @Tags ordens-de-servico
// @Accept json
// @Produce json
// @Param id path int true "ID da Ordem de Serviço"
// @Param ordem body dto.OrdemServicoCreateRequest true "Novos dados da Ordem de Serviço"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /ordens/{id} [put]
func (h *OrdemDeServicoHandler) UpdateOrdem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.OrdemServicoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	ordem := models.OrdemDeServico{
		DescricaoServico: req.DescricaoServico,
		Status:           req.Status,
		Responsavel:      req.Responsavel,
		OrcamentoID:      req.OrcamentoID,
	}

	if err := h.service.UpdateOrdem(id, ordem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar ordem de serviço"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ordem de serviço atualizada com sucesso"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteOrdem(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao excluir ordem de serviço"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ordem de serviço excluída com sucesso"})
}
