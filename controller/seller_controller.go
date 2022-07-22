package controller

import (
	 "TFP/domain/seller"
	 "TFP/utils"
	 "TFP/service"
	 "github.com/gin-gonic/gin"
	 "net/http"
	 "strings"
	 "TFP/helper"
)

var (
	SellersController SellersControllerInterface = &sellersController{}
)

type SellersControllerInterface interface {
 Create() gin.HandlerFunc
 GetAll() gin.HandlerFunc
 GetByID() gin.HandlerFunc
 Delete() gin.HandlerFunc
 Login() gin.HandlerFunc
 SignOut() gin.HandlerFunc
}

type sellersController struct{}

func (s *sellersController) Create() gin.HandlerFunc {
	return func (c *gin.Context) {
		var seller seller.Seller

		if err := c.BindJSON(&seller); err != nil {
			response := utils.BuildErrorResponse("Invalid JSON body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.SellersService.Create(seller)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		    response := utils.BuildResponse(result)
			c.JSON(http.StatusOK, response)
	}
}

func (s *sellersController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, getErr := service.SellersService.GetAll()
		if getErr != nil {
			utils.BuildErrorResponse(getErr)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *sellersController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("seller_id")
		if strings.TrimSpace(queryParam) == "" {
			utils.BuildErrorResponse("empty query param")
			return
		}

		user, getErr := service.SellersService.GetByID(queryParam)
		if getErr != nil {
			response := utils.BuildErrorResponse(getErr)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := utils.BuildResponse(user)
		c.JSON(http.StatusOK, response)
	}
}

func (s *sellersController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("seller_id")
		if strings.TrimSpace(queryParam) == "" {
			utils.BuildErrorResponse("empty query param")
			return
		}

		result, err := service.SellersService.Delete(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *sellersController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request seller.LoginRequest
		if err := c.BindJSON(&request); err != nil {
			response := utils.BuildErrorResponse("invalid json body")
			c.JSON(http.StatusNotFound, response)
			return
		}

		seller, err := service.SellersService.Login(request)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusNotFound, response)
			return
		}
		token, _ := helper.GenerateAlltokens(seller.Email, seller.FirstName, seller.LastName, seller.Sellerid)
		c.SetCookie("token", token, 24,"/", "", false, true)
		
		response := utils.BuildResponse(seller)
		c.JSON(http.StatusOK, response)
	}
}

func (s *sellersController) SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token","",-1,"", "",false,true)
		response := utils.BuildSignOutResponse()
		c.JSON(http.StatusOK, response)
	}
} 