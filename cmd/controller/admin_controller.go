package controller

import (
	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllAdmins(c *gin.Context) {
	var admins []models.Admin
	db.DB.Find(&admins)
	c.JSON(http.StatusOK, admins)
}

func GetAdminByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var admin models.Admin

	if err := db.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin não encontrado"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func CreateAdmin(c *gin.Context) {
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&admin)
	c.JSON(http.StatusCreated, admin)
}

func UpdateAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var admin models.Admin

	if err := db.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&admin)
	c.JSON(http.StatusOK, admin)
}

func DeleteAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var admin models.Admin

	if err := db.DB.First(&admin, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin não encontrado"})
		return
	}

	db.DB.Delete(&admin)
	c.Status(http.StatusNoContent)
}
