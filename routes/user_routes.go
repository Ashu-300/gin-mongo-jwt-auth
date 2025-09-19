package routes

import (
	"gin-jwt-auth/controllers"
	"gin-jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine){
	router.POST("/signup" , controllers.SignUp)
	router.POST("/login" , controllers.Login)
	

	
	
	router.Use( middleware.AuthMiddleware())
	router.GET("/profile"  , controllers.Profile)
	
	
}