package controller

import (
	"TFP/domain/buyer"
	"TFP/helper"
	"TFP/service"
	"TFP/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	UsersController UsersControllerInterface = &usersController{}
)

type UsersControllerInterface interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetByID() gin.HandlerFunc
	Delete() gin.HandlerFunc
	Login() gin.HandlerFunc
	SignOut() gin.HandlerFunc
	CheckUserExist() gin.HandlerFunc
}

type usersController struct{}

func (s *usersController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user buyer.User

		if err := c.BindJSON(&user); err != nil {
			response := utils.BuildErrorResponse("Invalid JSON body")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.UsersService.Create(user)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *usersController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, getErr := service.UsersService.GetAll()
		if getErr != nil {
			utils.BuildErrorResponse(getErr)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *usersController) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("userid")
		if strings.TrimSpace(queryParam) == "" {
			utils.BuildErrorResponse("empty query param")
			return
		}

		user, getErr := service.UsersService.GetByID(queryParam)
		if getErr != nil {
			response := utils.BuildErrorResponse(getErr)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := utils.BuildResponse(user)
		c.JSON(http.StatusOK, response)
	}
}

func (s *usersController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParam := c.Query("userid")
		if strings.TrimSpace(queryParam) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.UsersService.Delete(queryParam)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *usersController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request buyer.LoginRequest
		if err := c.BindJSON(&request); err != nil {
			response := utils.BuildErrorResponse("invalid json body")
			c.JSON(http.StatusNotFound, response)
			return
		}

		user, err := service.UsersService.Login(request)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusNotFound, response)
			return
		}

		token, _ := helper.GenerateAlltokens(user.Email, user.FirstName, user.LastName, user.Userid)

		c.SetCookie("token", token, 24, "/", "", false, true)

		response := utils.BuildResponse(user)
		c.JSON(http.StatusOK, response)
	}
}

func (s *usersController) SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("sign out application.......")
		c.SetCookie("token", "", -1, "", "", false, true)
		response := utils.BuildSignOutResponse()
		c.JSON(http.StatusOK, response)
	}
}

func (s *usersController) CheckUserExist() gin.HandlerFunc {
	return func(c *gin.Context) {
		emailID := c.Query("email_id")
		if strings.TrimSpace(emailID) == "" {
			response := utils.BuildErrorResponse("empty query param")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		result, err := service.UsersService.CheckUserExist(emailID)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *usersController) VerifyOTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request buyer.VerifyOTP
		if err := c.BindJSON(&request); err != nil {
			response := utils.BuildErrorResponse("invalid json body")
			c.JSON(http.StatusNotFound, response)
			return
		}

		result, err := service.UsersService.VerifyOTP(request)
		if err != nil {
			response := utils.BuildErrorResponse(err)
			c.JSON(http.StatusNotFound, response)
			return
		}

		response := utils.BuildResponse(result)
		c.JSON(http.StatusOK, response)
	}
}

func (s *usersController) UpdatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
