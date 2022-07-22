package controller

import (
	wishlist2 "TFP/domain/wishlist"
	"TFP/service"
	"TFP/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	WishlistController WishlistControllerInterface = &wishlistController{}
)

type WishlistControllerInterface interface {
	SaveToWishlist() gin.HandlerFunc
	GetWishlist() gin.HandlerFunc
	RemoveFromWishlist() gin.HandlerFunc
	ResetWishlist() gin.HandlerFunc
	MoveToCart() gin.HandlerFunc
}

type wishlistController struct{}

func (s *wishlistController) SaveToWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var wishlist wishlist2.Wishlist

		if err := c.ShouldBindJSON(&wishlist); err != nil {
			response := utils.BuildErrorResponse("Invalid JSON body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.WishlistService.Create(wishlist)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *wishlistController) GetWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("userID")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, totalPrice, err := service.WishlistService.GetUserWishlist(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildCartResponse(result, totalPrice)
		c.JSON(http.StatusOK, response)
	}
}

func (s *wishlistController) RemoveFromWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		if strings.TrimSpace(userID) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		productID := c.Query("productID")
		if strings.TrimSpace(productID) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.WishlistService.Remove(userID, productID)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *wishlistController) ResetWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		if strings.TrimSpace(userID) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.WishlistService.ResetWishlist(userID)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *wishlistController) MoveToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		if strings.TrimSpace(userID) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.WishlistService.MoveToCart(userID)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}
