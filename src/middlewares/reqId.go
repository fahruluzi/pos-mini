package middlewares

import (
	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/gin-gonic/gin"
)

//*Uses for unique each requested to server
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := utils.GenerateUuid()
		c.Writer.Header().Set("X-Request-Id", uuid)
		c.Next()
	}
}
