package routes

import (
	"TFP/controller"
	"github.com/gin-gonic/gin"
)

func WishlistRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/wishlist/create", controller.WishlistController.SaveToWishlist())
	incomingRoutes.GET("/wishlist/listwishlist", controller.WishlistController.GetWishlist())
	incomingRoutes.DELETE("/wishlist/remove", controller.WishlistController.RemoveFromWishlist())
	incomingRoutes.DELETE("/wishlist/resetwishlist", controller.WishlistController.ResetWishlist())
	incomingRoutes.POST("/wishlist/movetocart", controller.WishlistController.MoveToCart())
}
