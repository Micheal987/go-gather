package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// home
func UserHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": struct{ name string }{name: "hd"},
		"id":   c.Params.ByName("id"),
	})
}

func UserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "userinfo",
	})
}

// add
