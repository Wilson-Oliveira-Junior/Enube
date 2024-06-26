package routes

import (
	"myapi/controllers"
	"myapi/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.GET("/data", middleware.AuthMiddleware(), controllers.GetData)
	r.POST("/upload", controllers.UploadFile)

	return r
}
