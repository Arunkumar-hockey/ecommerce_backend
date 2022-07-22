package controller

import (
	"net/http"
	"TFP/domain/category"
	"github.com/gin-gonic/gin"
	"TFP/utils"
	"TFP/service"
)

var (
	CategoryController CategoryControllerInterface = &categoryController{}
)


type CategoryControllerInterface interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
}

type categoryController struct {}

func (s *categoryController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category category.Category

		if err := c.BindJSON(&category); err != nil {
			response := utils.BuildErrorResponse("invalid json body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.CategoryService.Create(category)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *categoryController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, getErr := service.CategoryService.GetAll()
		if getErr != nil {
			response := utils.BuildErrorResponse(getErr)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}