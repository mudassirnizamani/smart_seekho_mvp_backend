package main

import (
	"os"
	"smart_seekho_mvp/src/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var port string = os.Getenv("port")

	if port == "" {
		port = "4000"
	}

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	// Pipeline
	config.AuthRoutes(app)

	app.Run(":" + port)
}
