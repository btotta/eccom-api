package middleware

import (
	common_error "eccom-api/internal/pkg/erros"
	"eccom-api/internal/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		if c.Request.Header.Get("Authorization") == "" {
			common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "Authorization header não informado")
			c.Abort()
			return
		}

		auth := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(auth) != 2 {
			common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "Authorization header inválido")
			c.Abort()
			return
		}

		if auth[0] != "Bearer" {
			common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "Authorization header inválido")
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(auth[1])
		if err != nil {
			common_error.DefaultErrorResponse(c, http.StatusUnauthorized, "Token inválido")
			c.Abort()
			return
		}

		c.Set("email", claims.Email)

		c.Next()
	}
}
