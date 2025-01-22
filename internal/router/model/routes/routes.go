package routes

import (
	"gintest/internal/router/model/routes/server"

	"github.com/gin-gonic/gin"
)

func RouteDefault(r *gin.Engine) {
	r.GET("/", server.DefaultHome)
}
