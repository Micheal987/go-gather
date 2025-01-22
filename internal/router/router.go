package router

import (
	"gintest/internal/router/model/handler"
	"gintest/internal/router/model/routes"
	"gintest/internal/router/model/routes/auth"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
	Emial    string `form:"email" binding:"required"`
}

func MainRouter() {
	//默认路由
	r := gin.New()
	//自定義中間件
	r.Use(handler.Logger())
	routes.RouteDefault(r)
	auth.RoutesAuth(r)

	r.Run(":8000")
}
