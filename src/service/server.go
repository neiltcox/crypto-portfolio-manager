package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()

	// All static files are accessible at `/static`
	router.Static("/static", "../frontend/dist/")

	// The index file is accessible at the root
	router.StaticFile("/", "../frontend/dist/index.html")

	apiV1 := router.Group("/api/v1")
	apiV1.Use(Authenticate())

	apiV1.GET("/strategy", GetStrategy())
	apiV1.POST("/strategy", PostStrategy())
	apiV1.GET("/portfolio", GetPortfolio())
	apiV1.GET("/portfolios", GetPortfolios())
	apiV1.POST("/portfolio", PostPortfolio())
	apiV1.GET("/user", GetUser())

	http.ListenAndServe(":5010", router)
}
