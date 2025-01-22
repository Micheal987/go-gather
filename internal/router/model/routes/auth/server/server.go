package server

import (
	"gintest/internal/db/collections/create"
	"gintest/internal/db/instance"
	"gintest/internal/db/model/user"
	"gintest/internal/jwtCombined"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	status := http.StatusOK
	if name == "" {
		ctx.JSON(500, gin.H{"msg": "error"})
		return
	}
	ctx.JSON(status, gin.H{
		"Status": status,
		"data":   saveUserSql(name, password, email),
	})
}
func LoginUser(ctx *gin.Context) {
	token := ctx.MustGet("Token").(string)
	jwtCombined.CreateToken()
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// sql
func saveUserSql(name string, password string, email string) *user.User {
	result := user.User{Name: name, Password: &password, Email: &email}
	info := create.CreateSingle(instance.DbInstance, instance.ErrorInstance, &result)
	return info
}
