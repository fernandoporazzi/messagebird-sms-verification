package main

import (
	"os"

	"github.com/fernandoporazzi/messagebird-sms-verification/apis"
	"github.com/fernandoporazzi/messagebird-sms-verification/config"
	"github.com/gin-gonic/gin"
	messagebird "github.com/messagebird/go-rest-api"
)

func main() {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", apis.Health)
		v1.POST("/create", apis.Create)
		v1.GET("/verify/:id/:token", apis.Verify)
	}

	config.Config.Client = messagebird.New(os.Getenv("ACCESS_KEY"))

	r.Run()
}
