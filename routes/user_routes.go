package routes

import (
	"TFP/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/buyer/create", controller.UsersController.Create())
	incomingRoutes.POST("/buyer/login", controller.UsersController.Login())
	incomingRoutes.POST("/buyer/checkuserexist", controller.UsersController.CheckUserExist())
}

func SecuredUserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/buyer/getall", controller.UsersController.GetAll())
	incomingRoutes.GET("/buyer/getuserbyid", controller.UsersController.GetByID())
	incomingRoutes.DELETE("/buyer/delete", controller.UsersController.Delete())
	incomingRoutes.POST("/buyer/signout", controller.UsersController.SignOut())
}
