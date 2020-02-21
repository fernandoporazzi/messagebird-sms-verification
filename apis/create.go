package apis

import (
	"log"
	"net/http"

	"github.com/fernandoporazzi/messagebird-sms-verification/config"
	"github.com/fernandoporazzi/messagebird-sms-verification/models"
	"github.com/gin-gonic/gin"
	"github.com/messagebird/go-rest-api/verify"
)

// Create fetchs a new Verify object from Messagebird and returns the whole payload to the client
func Create(c *gin.Context) {
	var input models.Input

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	v, err := verify.Create(config.Config.Client, input.Phone, &verify.Params{
		Originator: "Fernando",
	})

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}

	c.JSON(http.StatusCreated, v)
}
