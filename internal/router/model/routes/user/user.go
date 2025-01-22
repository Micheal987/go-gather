package user

import "github.com/gin-gonic/gin"

func userRoutes(r *gin.Engine) {
	r.Group("/user")

}
