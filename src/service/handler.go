package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStrategy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func PostStrategy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func GetPortfolio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func GetPortfolios() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUserId := DistillAuthUserId(ctx)
		portfolios := FindPortfoliosByUserId(authUserId)

		buildStandardResponse(
			ctx,
			gin.H{
				"Portfolios": portfolios,
			},
		)
	}
}

func PostPortfolio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func buildStandardResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
