package user

import (
	"net/http"

	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	response := utils.Response{C: c}
	registerValidator := NewRegisterValidator()

	if err := registerValidator.Bind(c); err != nil {
		response.ResponseFormatter(http.StatusNotAcceptable, "Invalid Form", err, gin.H{"err_message": err.Error()})
		return
	}

	response.ResponseFormatter(200, "hihihi", nil, nil)
}
