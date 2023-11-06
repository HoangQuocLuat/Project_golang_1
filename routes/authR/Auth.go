package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	userRoutes := r.Group("auth")
	{
		userRoutes.POST("/register", handler.RegisterHandler())
	}
}
