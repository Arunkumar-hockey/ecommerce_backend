package routes

import (
	"TFP/controller"
	"github.com/gin-gonic/gin"
)

func BuyRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/order/buyfromcart", controller.OrderController.PlaceOrder())
	incomingRoutes.GET("/order/getorderbyuser", controller.OrderController.GetOrderByUser())
	incomingRoutes.GET("/order/getallorders", controller.OrderController.GetAllOrder())

}
