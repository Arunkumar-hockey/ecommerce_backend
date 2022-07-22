package routes

import (
	"github.com/gin-gonic/gin"
	"TFP/controller"
)

func SellerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/seller/create", controller.SellersController.Create())
	incomingRoutes.POST("/seller/login", controller.SellersController.Login())
}

func SecuredSellerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/seller/getall", controller.SellersController.GetAll())
	incomingRoutes.GET("/seller/getsellerbyid", controller.SellersController.GetByID())
	incomingRoutes.DELETE("/seller/delete", controller.SellersController.Delete())
	incomingRoutes.POST("/seller/signout", controller.SellersController.SignOut())
}