package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func ResponseCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}
