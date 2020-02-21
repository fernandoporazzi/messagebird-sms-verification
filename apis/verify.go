package apis

import (
	"log"
	"net/http"

	"github.com/fernandoporazzi/messagebird-sms-verification/config"
	"github.com/gin-gonic/gin"
	"github.com/messagebird/go-rest-api/verify"
)

// Verify verifies a sent verification token. Can only be done once for each token.
func Verify(c *gin.Context) {
	id := c.Param("id")
	token := c.Param("token")

	v, err := verify.VerifyToken(
		config.Config.Client,
		id,
		token,
	)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	c.JSON(http.StatusCreated, v)
}
