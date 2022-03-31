package routes

import (
	"boilerplate/modules/auth"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(route *gin.Engine) {

	/**
	@description All Handler Auth
	*/

	authService := auth.NewServiceRegister()

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/auth")
	groupRoute.POST("/register", authService.Register)
	groupRoute.POST("/login", authService.Login)
}
