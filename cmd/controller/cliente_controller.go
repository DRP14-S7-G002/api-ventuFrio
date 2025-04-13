package controller

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllClientes(c *gin.Context) {
	var clientes []models.Cliente
	db.DB.Find(&clientes)
	c.JSON(http.StatusOK, clientes)
}

func GetClienteByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cliente models.Cliente

	if err := db.DB.First(&cliente, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}
	c.JSON(http.StatusOK, cliente)
}

func CreateCliente(c *gin.Context) {
	var cliente models.Cliente

	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&cliente)
	c.JSON(http.StatusCreated, cliente)
}

func UpdateCliente(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cliente models.Cliente

	if err := db.DB.First(&cliente, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&cliente)
	c.JSON(http.StatusOK, cliente)
}

func DeleteCliente(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cliente models.Cliente

	if err := db.DB.First(&cliente, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	db.DB.Delete(&cliente)
	c.Status(http.StatusNoContent)
}
