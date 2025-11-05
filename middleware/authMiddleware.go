package middleware

import (
	"net/http"
	"strings"

	"go-task-manager-mvc/config"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware protege rutas que requieren autenticación JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Falta el token de autorización"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := config.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		// Guardar usuario en el contexto para usarlo en controladores
		c.Set("username", claims.Username)
		c.Next()
	}
}
