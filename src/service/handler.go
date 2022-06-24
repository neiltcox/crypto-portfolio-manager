package service

import (
	"net/http"
	"strconv"

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
		authUserId := DistillAuthUserId(ctx)

		portfolioId, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			// TODO: error
		}

		portfolio, err := FindPortfolioById(uint(portfolioId))
		if err != nil {
			// TODO: error
		}

		if authUserId != uint(portfolio.UserID) {
			// TODO: error
		}

		buildStandardResponse(
			ctx,
			gin.H{
				"Portfolio": portfolio,
			},
		)
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

func GetExchangeSupportedAssets() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		portfolioId, err := strconv.Atoi(ctx.Query("portfolio_id"))
		if err != nil {
			// TODO: error
		}

		portfolio, err := FindPortfolioById(uint(portfolioId))
		if err != nil {
			// TODO: error
		}

		exchange, err := portfolio.Exchange()
		if err != nil {
			// TODO: error
		}

		supportedAssets, err := exchange.SupportedAssets(portfolio)
		if err != nil {
			// TODO: error
		}

		ctx.JSON(http.StatusOK, gin.H{
			"ExchangeSupportedAssets": supportedAssets,
		})
	}
}

func GetExchangeConnectionValid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUserId := DistillAuthUserId(ctx)

		portfolioId, err := strconv.Atoi(ctx.Query("portfolio_id"))
		if err != nil {
			// TODO: error
		}

		portfolio, err := FindPortfolioById(uint(portfolioId))
		if err != nil {
			// TODO: error
		}

		if authUserId != uint(portfolio.UserID) {
			// TODO: error
		}

		exchange, err := portfolio.Exchange()
		if err != nil {
			// TODO: error
		}

		testExchangeConnectionResult := exchange.ValidateConnection(portfolio)

		buildStandardResponse(
			ctx,
			gin.H{
				"TestExchangeConnectionResult": testExchangeConnectionResult,
			},
		)
	}
}

func buildStandardResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}
