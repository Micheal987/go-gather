package auth

import (
	"gintest/internal/router/model/handler"
	"gintest/internal/router/model/routes/auth/server"

	"github.com/gin-gonic/gin"
)

func RoutesAuth(r *gin.Engine) {
	auth := r.Group("/auth").Use(handler.Logger())
	auth.POST("/register", server.RegisterUser)
	auth.POST("/login", server.LoginUser)
}
