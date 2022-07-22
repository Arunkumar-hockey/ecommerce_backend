package routes

import (
	"TFP/controller"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/admin/create", controller.AdminController.Create())
	incomingRoutes.POST("/admin/login", controller.AdminController.Login())
}

func SecuredAdminRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/admin/getall", controller.AdminController.GetAll())
	incomingRoutes.GET("/admin/getadminbyid", controller.AdminController.GetByID())
	incomingRoutes.DELETE("/admin/delete", controller.AdminController.Delete())
	incomingRoutes.POST("/admin/signout", controller.AdminController.SignOut())
}
