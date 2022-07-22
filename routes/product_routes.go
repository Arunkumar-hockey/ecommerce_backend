package routes

import (
	"TFP/controller"
	"github.com/gin-gonic/gin"
)

func ProductsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/products/create", controller.ProductsController.Create())
	incomingRoutes.GET("/products/getall", controller.UsersController.GetAll())
	incomingRoutes.GET("/products/getproductbyid", controller.ProductsController.GetByID())
	incomingRoutes.GET("products/getproductbycategory", controller.ProductsController.GetProductByCategory())
	incomingRoutes.PATCH("/products/update", controller.ProductsController.Update())
	incomingRoutes.DELETE("/products/delete", controller.ProductsController.Delete())
	incomingRoutes.GET("/redis/getbyid", controller.ProductsController.GetRedisData())
	incomingRoutes.GET("/redis/getall", controller.ProductsController.GetAllRedisData())
	incomingRoutes.DELETE("redis/deletebyid", controller.ProductsController.DeleteRedisData())
}
