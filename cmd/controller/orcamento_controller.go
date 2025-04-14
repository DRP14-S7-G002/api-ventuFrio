package controller

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllOrcamentos(c *gin.Context) {
	var orcamentos []models.Orcamento
	db.DB.Find(&orcamentos)
	c.JSON(http.StatusOK, orcamentos)
}

func GetOrcamentoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var orcamento models.Orcamento

	if err := db.DB.First(&orcamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orçamento não encontrado"})
		return
	}
	c.JSON(http.StatusOK, orcamento)
}

func CreateOrcamento(c *gin.Context) {
	var orcamento models.Orcamento

	if err := c.ShouldBindJSON(&orcamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&orcamento)
	c.JSON(http.StatusCreated, orcamento)
}

func UpdateOrcamento(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var orcamento models.Orcamento

	if err := db.DB.First(&orcamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orçamento não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&orcamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&orcamento)
	c.JSON(http.StatusOK, orcamento)
}

func DeleteOrcamento(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var orcamento models.Orcamento

	if err := db.DB.First(&orcamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Orçamento não encontrado"})
		return
	}

	db.DB.Delete(&orcamento)
	c.Status(http.StatusNoContent)
}
