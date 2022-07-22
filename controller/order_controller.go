package controller

import (
	"TFP/service"
	"TFP/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	OrderController OrderControllerInterface = &orderController{}
)

type OrderControllerInterface interface {
	PlaceOrder() gin.HandlerFunc
	GetOrderByUser() gin.HandlerFunc
	GetAllOrder() gin.HandlerFunc
}

type orderController struct{}

func (s *orderController) PlaceOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		if strings.TrimSpace(userID) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.OrderService.PlaceOrder(userID)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *orderController) GetOrderByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("userID")
		if strings.TrimSpace(userID) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.OrderService.GetOrder(userID)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *orderController) GetAllOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := service.OrderService.GetAllOrder()
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}
