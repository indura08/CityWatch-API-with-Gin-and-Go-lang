package routes

import (
	"citywatch/internal/handler"
	"citywatch/internal/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, userHandler *handler.UserHandler) {

	auth := r.Group("/api/auth/")

	auth.POST("/register/admin", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0), userHandler.AdminRegistration)
	auth.POST("/register/worker", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0), userHandler.WorkerRegister)
	auth.POST("/register/citizen", userHandler.CitizenRegister)

	auth.POST("/login", userHandler.Login)
	auth.GET("/hi/test", middleware.AuthorizeRoles(os.Getenv("SECRET"), 0, 1), func(c *gin.Context) {
		c.JSON(201, gin.H{
			"message": "hello to citizen or admin",
		})
	})

}
