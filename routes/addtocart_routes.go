package routes

import (
	"TFP/controller"
	"github.com/gin-gonic/gin"
)

func AddToCartRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/addtocart/create", controller.AddToCartController.SaveToCart())
	incomingRoutes.GET("/addtocart/listcart", controller.AddToCartController.GetUserCart())
	incomingRoutes.DELETE("/addtocart/remove", controller.AddToCartController.RemoveFromCart())
	incomingRoutes.DELETE("/addtocart/resetcart", controller.AddToCartController.ResetCart())
}
