package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//reqAPIKey := ctx.Request.URL.User.Username()
	}
}
