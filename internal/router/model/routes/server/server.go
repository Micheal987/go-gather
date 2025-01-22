package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultHome(c *gin.Context) {
	code := http.StatusOK
	c.JSON(code, gin.H{
		"code": code,
		"data": map[string][]string{"relust": {""}},
	})
}
