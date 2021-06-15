package main

import (
	"net/http"

	"github.com/fahruluzi/pos-mini/src/middlewares"
	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middlewares.RequestIDMiddleware())

	r.NoRoute(func(c *gin.Context) {
		utils.ResponseFormatter(http.StatusNotFound, "Not Found.", nil, nil, c)
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			utils.ResponseFormatter(http.StatusOK, "Pong", nil, nil, c)
		})
	}

	return r
}
