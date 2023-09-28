package config

import (
	"smart_seekho_mvp/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(app *gin.Engine) {
	app.POST("/generate-opt", controllers.GenerateOtp())
	app.POST("/authenticate", controllers.Authenticate())
}
