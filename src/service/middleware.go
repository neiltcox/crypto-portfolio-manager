package service

import "github.com/gin-gonic/gin"

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}
