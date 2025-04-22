package handler

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/service"
	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	service service.MaterialService
}

func NewMaterialHandler(s service.MaterialService) *MaterialHandler {
	return &MaterialHandler{s}
}

// GetAllMateriais godoc
// @Summary Lista todos os materiais
// @Description Retorna todos os registros de materiais
// @Tags materiais
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Material
// @Failure 500 {object} map[string]string
// @Router /materials [get]
func (h *MaterialHandler) GetAllMateriais(c *gin.Context) {
	materiais, err := h.service.GetAllMateriais()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving materiais"})
		return
	}
	c.JSON(http.StatusOK, materiais)
}

// GetMaterialByID godoc
// @Summary Busca um material pelo ID
// @Description Retorna um material com base no ID fornecido
// @Tags materiais
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Material"
// @Success 200 {object} models.Material
// @Failure 400,404 {object} map[string]string
// @Router /materials/{id} [get]
func (h *MaterialHandler) GetMaterialByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	material, err := h.service.GetMaterialByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}
	c.JSON(http.StatusOK, material)
}

// CreateMaterial godoc
// @Summary Cria um novo material
// @Description Registra um novo material com os dados fornecidos
// @Tags materiais
// @Accept  json
// @Produce  json
// @Param material body models.Material true "Dados do Material"
// @Success 201 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /materials [post]
func (h *MaterialHandler) CreateMaterial(c *gin.Context) {
	var material models.Material
	if err := c.ShouldBindJSON(&material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.CreateMaterial(material); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create material"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Material created successfully"})
}

// UpdateMaterial godoc
// @Summary Atualiza um material existente
// @Description Atualiza os dados de um material pelo ID
// @Tags materiais
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Material"
// @Param material body models.Material true "Novos dados do Material"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /materials/{id} [put]
func (h *MaterialHandler) UpdateMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var material models.Material
	if err := c.ShouldBindJSON(&material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.UpdateMaterial(id, material); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Material updated"})
}

// DeleteMaterial godoc
// @Summary Remove um material
// @Description Exclui um material com base no ID fornecido
// @Tags materiais
// @Accept  json
// @Produce  json
// @Param id path int true "ID do Material"
// @Success 200 {object} map[string]string
// @Failure 400,500 {object} map[string]string
// @Router /materials/{id} [delete]
func (h *MaterialHandler) DeleteMaterial(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteMaterial(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Material deleted"})
}
