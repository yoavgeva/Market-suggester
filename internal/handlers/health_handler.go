package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}