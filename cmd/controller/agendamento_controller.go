package controller

import (
	"github.com/DRP14-S7-G002/api-ventuFrio/internal/db"

	"net/http"
	"strconv"

	"github.com/DRP14-S7-G002/api-ventuFrio/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllAgendamentos(c *gin.Context) {
	var agendamentos []models.Agendamento
	db.DB.Find(&agendamentos)
	c.JSON(http.StatusOK, agendamentos)
}

func GetAgendamentoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var agendamento models.Agendamento

	if err := db.DB.First(&agendamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}
	c.JSON(http.StatusOK, agendamento)
}

func CreateAgendamento(c *gin.Context) {
	var agendamento models.Agendamento

	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&agendamento)
	c.JSON(http.StatusCreated, agendamento)
}

func UpdateAgendamento(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var agendamento models.Agendamento

	if err := db.DB.First(&agendamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&agendamento)
	c.JSON(http.StatusOK, agendamento)
}

func DeleteAgendamento(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var agendamento models.Agendamento

	if err := db.DB.First(&agendamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	db.DB.Delete(&agendamento)
	c.Status(http.StatusNoContent)
}
