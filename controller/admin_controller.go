package controller

import (
	"TFP/domain/admin"
	"TFP/helper"
	"TFP/service"
	"TFP/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	AdminController AdminControllerInterface = &adminController{}
)

type AdminControllerInterface interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	Delete() gin.HandlerFunc
	Login() gin.HandlerFunc
	SignOut() gin.HandlerFunc
}

type adminController struct{}

func (s *adminController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var admin admin.Admin

		if err := c.BindJSON(&admin); err != nil {
			response := utils.BuildErrorResponse("Invalid JSON body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.AdminService.Create(admin)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *adminController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, getErr := service.AdminService.GetAll()
		if getErr != nil {
			utils.BuildErrorResponse(getErr)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *adminController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("adminId")
		if strings.TrimSpace(queryParam) == "" {
			utils.BuildErrorResponse("empty query param")
			return
		}

		user, getErr := service.AdminService.GetByID(queryParam)
		if getErr != nil {
			response := utils.BuildErrorResponse(getErr)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := utils.BuildResponse(user)
		c.JSON(http.StatusOK, response)
	}
}

func (s *adminController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("adminId")
		if strings.TrimSpace(queryParam) == "" {
			utils.BuildErrorResponse("empty query param")
			return
		}

		result, err := service.AdminService.Delete(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *adminController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request admin.LoginRequest
		if err := c.BindJSON(&request); err != nil {
			response := utils.BuildErrorResponse("invalid json body")
			c.JSON(http.StatusNotFound, response)
			return
		}

		admin, err := service.AdminService.Login(request)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusNotFound, response)
			return
		}

		token, _ := helper.GenerateAlltokens(admin.Email, admin.FirstName, admin.LastName, admin.AdminID)

		c.SetCookie("token", token, 24, "/", "", false, true)

		response := utils.BuildResponse(admin)
		c.JSON(http.StatusOK, response)
	}
}

func (s *adminController) SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "", "", false, true)
		response := utils.BuildSignOutResponse()
		c.JSON(http.StatusOK, response)
	}
}
