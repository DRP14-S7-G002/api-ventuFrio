package controller

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllEnderecos(c *gin.Context) {
	var enderecos []models.Endereco
	db.DB.Find(&enderecos)
	c.JSON(http.StatusOK, enderecos)
}

func GetEnderecoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var endereco models.Endereco

	if err := db.DB.First(&endereco, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endereço não encontrado"})
		return
	}
	c.JSON(http.StatusOK, endereco)
}

func CreateEndereco(c *gin.Context) {
	var endereco models.Endereco

	if err := c.ShouldBindJSON(&endereco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&endereco)
	c.JSON(http.StatusCreated, endereco)
}

func UpdateEndereco(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var endereco models.Endereco

	if err := db.DB.First(&endereco, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endereço não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&endereco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&endereco)
	c.JSON(http.StatusOK, endereco)
}

func DeleteEndereco(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var endereco models.Endereco

	if err := db.DB.First(&endereco, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endereço não encontrado"})
		return
	}

	db.DB.Delete(&endereco)
	c.Status(http.StatusNoContent)
}
