package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/dto"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	service service.MaterialService
}

func NewMaterialHandler(s service.MaterialService) *MaterialHandler {
	return &MaterialHandler{service: s}
}

// GetAllMateriais godoc
// @Summary Lista todos os materiais
// @Description Retorna todos os registros de materiais
// @Tags materiais
// @Accept json
// @Produce json
// @Success 200 {array} dto.MaterialResponse
// @Failure 500 {object} map[string]string
// @Router /materials [get]
func (h *MaterialHandler) GetAllMateriais(c *gin.Context) {
	materiais, err := h.service.GetAllMateriais()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar materiais"})
		return
	}

	var response []dto.MaterialResponse
	for _, material := range materiais {
		response = append(response, dto.MaterialResponse{
			ID:               material.ID,
			Nome:             material.Nome,
			Quantidade:       material.Quantidade,
			Valor:            material.Valor,
			OrdemDeServicoID: material.OrdemDeServicoID,
		})
	}

	c.JSON(http.StatusOK, response)
}

// GetMaterialByID godoc
// @Summary Busca um material pelo ID
// @Description Retorna um material com base no ID fornecido
// @Tags materiais
// @Accept json
// @Produce json
// @Param id path int true "ID do Material"
// @Success 200 {object} dto.MaterialResponse
// @Failure 400,404 {object} map[string]string
// @Router /materials/{id} [get]
func (h *MaterialHandler) GetMaterialByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	material, err := h.service.GetMaterialByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material não encontrado"})
		return
	}

	response := dto.MaterialResponse{
		ID:               material.ID,
		Nome:             material.Nome,
		Quantidade:       material.Quantidade,
		Valor:            material.Valor,
		OrdemDeServicoID: material.OrdemDeServicoID,
	}

	c.JSON(http.StatusOK, response)
}

// CreateMaterial godoc
// @Summary Cria um novo material
// @Description Registra um novo material com os dados fornecidos
// @Tags materiais
// @Accept json
// @Produce json
// @Param material body dto.MaterialCreateRequest true "Dados do Material"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /materials [post]
func (h *MaterialHandler) CreateMaterial(c *gin.Context) {
	var req dto.MaterialCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	material := models.Material{
		Nome:             req.Nome,
		Quantidade:       req.Quantidade,
		Valor:            req.Valor,
		OrdemDeServicoID: req.OrdemDeServicoID,
	}

	if err := h.service.CreateMaterial(material); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar material"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Material criado com sucesso"})
}

// UpdateMaterial godoc
// @Summary Atualiza um material existente
// @Description Atualiza os dados de um material pelo ID
// @Tags materiais
// @Accept json
// @Produce json
// @Param id path int true "ID do Material"
// @Param material body dto.MaterialCreateRequest true "Novos dados do Material"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /materials/{id} [put]
func (h *MaterialHandler) UpdateMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req dto.MaterialCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entrada inválida"})
		return
	}

	material := models.Material{
		Nome:             req.Nome,
		Quantidade:       req.Quantidade,
		Valor:            req.Valor,
		OrdemDeServicoID: req.OrdemDeServicoID,
	}

	if err := h.service.UpdateMaterial(id, material); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao atualizar material"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Material atualizado com sucesso"})
}

// DeleteMaterial godoc
// @Summary Remove um material
// @Description Exclui um material com base no ID fornecido
// @Tags materiais
// @Accept json
// @Produce json
// @Param id path int true "ID do Material"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /materials/{id} [delete]
func (h *MaterialHandler) DeleteMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.service.DeleteMaterial(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao excluir material"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Material excluído com sucesso"})
}
