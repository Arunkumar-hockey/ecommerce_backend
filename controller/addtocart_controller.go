package controller

import (
	"TFP/domain/addtocart"
	"TFP/service"
	"TFP/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	AddToCartController AddToCartInterface = &addToCartController{}
)

type AddToCartInterface interface {
	SaveToCart() gin.HandlerFunc
	GetUserCart() gin.HandlerFunc
	RemoveFromCart() gin.HandlerFunc
	ResetCart() gin.HandlerFunc
}

type addToCartController struct{}

func (s *addToCartController) SaveToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cart addtocart.AddToCart

		if err := c.BindJSON(&cart); err != nil {
			response := utils.BuildErrorResponse("Invalid JSON body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.AddToCartService.Create(cart)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *addToCartController) GetUserCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("userID")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, totalPrice, err := service.AddToCartService.GetUserCart(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildCartResponse(result, totalPrice)
		c.JSON(http.StatusOK, response)
	}
}

func (s *addToCartController) RemoveFromCart() gin.HandlerFunc {
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

		result, err := service.AddToCartService.Remove(userID, productID)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *addToCartController) ResetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		if strings.TrimSpace(userID) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.AddToCartService.ResetCart(userID)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}
