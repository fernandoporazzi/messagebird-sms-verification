package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health retrieves the API status
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
