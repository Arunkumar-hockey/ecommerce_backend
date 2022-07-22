package routes

import (
	"github.com/gin-gonic/gin"
	"TFP/controller"
)

func CategoryRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/category/create", controller.CategoryController.Create())
	incomingRoutes.GET("/category/getall", controller.CategoryController.GetAll())
}
