package main

import (
	"net/http"

	"github.com/fahruluzi/pos-mini/docs"
	"github.com/fahruluzi/pos-mini/src/apps/user"
	"github.com/fahruluzi/pos-mini/src/middlewares"
	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// * PingApp godoc
// @Summary Pinging App
// @Description ping app
// @Tags Default
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.ResponseModel
// @Router /ping [get]
func Router() *gin.Engine {
	docs.SwaggerInfo.Title = "POS Mini"
	docs.SwaggerInfo.Description = "Aplikasi POS mini"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:15061"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middlewares.RequestIDMiddleware())
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.NoRoute(func(c *gin.Context) {
		response := utils.Response{C: c}
		response.ResponseFormatter(http.StatusNotFound, "Not Found.", nil, nil)
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api")
	{
		v1.GET("/ping", func(c *gin.Context) {
			response := utils.Response{C: c}
			response.ResponseFormatter(http.StatusOK, "Pong", nil, nil)
		})
	}
	v1.Use(middlewares.JWTAuthMiddleware(false))
	user.UserAnonymusRouter(v1.Group("/user"))

	v1.Use(middlewares.JWTAuthMiddleware(true))
	user.UserRouter(v1.Group("/user"))

	return r
}
