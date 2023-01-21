package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/signup", controllers.SignUp())
	incomingRoutes.POST("/api/login", controllers.Login())
	incomingRoutes.GET("/api/user", controllers.User())
	incomingRoutes.POST("/api/logout", controllers.Logout())
}
