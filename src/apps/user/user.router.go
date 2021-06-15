package user

import "github.com/gin-gonic/gin"

func UserAnonymusRouter(router *gin.RouterGroup) {
	router.POST("/register", UserCreate)
	router.POST("/login", UserLogin)
}
