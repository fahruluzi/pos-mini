package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := utils.Response{C: c}
		c.Set("my_user_model", 0)
		authHeader := c.GetHeader("access_token")
		if auto401 {
			if len(authHeader) <= 0 {
				response.ResponseFormatter(http.StatusUnauthorized, "Invalid Access Token. Re-login now", nil, nil)
				c.Abort()
				return
			}

			token, err := utils.ValidateToken(authHeader)
			if err != nil {
				response.ResponseFormatter(http.StatusUnauthorized, "Invalid Access Token. Re-login now", err, nil)
				c.Abort()
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			c.Set("my_user_model", claims)
		}

		c.Next()
	}
}
