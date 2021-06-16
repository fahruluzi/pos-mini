package product

import "github.com/gin-gonic/gin"

func ProductRouter(router *gin.RouterGroup) {
	router.GET("/", GetListProduct)
	router.GET("/:uuid", GetProductDetail)
	router.POST("/", CreateProduct)
	router.PATCH("/:uuid", UpdateProduct)
	router.DELETE("/:uuid", DeleteProduct)
}
