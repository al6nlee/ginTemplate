package routes

import (
	"ginTemplate/pkg/api/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/user/:id", controllers.GetUser)
}
