package handler

import (
	"gintest/internal/core"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Token", "admin")
		core.Zap("http://localhost:8000")
		ctx.Next() //放行
	}
}
