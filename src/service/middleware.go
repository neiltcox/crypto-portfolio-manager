package service

import "github.com/gin-gonic/gin"

const AuthUserIdKey = "auth_user_id"

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: don't hardcode, authenticate
		ctx.Set(AuthUserIdKey, 1)

		ctx.Next()
	}
}

func DistillAuthUserId(ctx *gin.Context) uint {
	return uint(ctx.GetInt(AuthUserIdKey))
}
