package user

import "github.com/gin-gonic/gin"

// * Register godoc
// @Summary Register User
// @Description register app
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param user body RegisterValidator true "Register User"
// @Success 200 {object} utils.ResponseModel
// @Router /user/register [post]
func UserAnonymusRouter(router *gin.RouterGroup) {
	router.POST("/register", UserCreate)
}
