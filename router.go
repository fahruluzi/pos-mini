package main

import (
	"net/http"

	"github.com/fahruluzi/pos-mini/docs"
	"github.com/fahruluzi/pos-mini/src/middlewares"
	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name API Support
// @contact.email fahrul.fauz@gmail.com
// @termsOfService http://swagger.io/terms/

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

	r.NoRoute(func(c *gin.Context) {
		utils.ResponseFormatter(http.StatusNotFound, "Not Found.", nil, nil, c)
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api")
	{
		v1.GET("/ping", func(c *gin.Context) {
			utils.ResponseFormatter(http.StatusOK, "Pong", nil, nil, c)
		})
	}

	return r
}
