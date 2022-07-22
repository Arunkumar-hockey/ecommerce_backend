package controller

import (
	"TFP/domain/products"
	"TFP/service"
	"TFP/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	ProductsController productsControllerInterface = &productsController{}
)

type productsControllerInterface interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	GetProductByCategory() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetRedisData() gin.HandlerFunc
	GetAllRedisData() gin.HandlerFunc
	DeleteRedisData() gin.HandlerFunc
}

type productsController struct{}

func (s *productsController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product products.Product
		if err := c.BindJSON(&product); err != nil {
			response := utils.BuildErrorResponse("Invalid JSON body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.ProductsService.Create(product)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *productsController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := service.ProductsService.GetAll()
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *productsController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("productid")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.ProductsService.GetByID(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *productsController) GetProductByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("categoryid")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.ProductsService.GetProductByCategory(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *productsController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product products.Product

		queryParam := c.Query("productid")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		isPartial := c.Request.Method == http.MethodPatch
		product.ProductID = queryParam
		result, err := service.ProductsService.Update(isPartial, product)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *productsController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("productid")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.ProductsService.Delete(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *productsController) GetRedisData() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("productid")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.ProductsService.GetRedisData(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *productsController) GetAllRedisData() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := service.ProductsService.GetAllRedisData()
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *productsController) DeleteRedisData() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("productid")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		result, err := service.ProductsService.DeleteRedisData(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}

}
