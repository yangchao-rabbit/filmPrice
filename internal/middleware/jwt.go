package middleware

import (
	"filmPrice/internal/apps/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var skipPaths = map[string]bool{
	"/api/auth/local-login": true,
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		fmt.Println(path)
		if path == "" {
			c.Next()
			return
		}

		if skipPaths[path] {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少或非法Token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token无效或已过期"})
			c.Abort()
			return
		}

		// 存入上下文
		c.Set("user", claims.User)
		c.Next()
	}
}
