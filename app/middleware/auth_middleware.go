package middleware

import (
	"context"
	"fmt"
	"net/http"

	http2 "iflow-lite/core/http"
	token2 "iflow-lite/core/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if tokenString != "" {
			token, err := token2.ValidateJWT(tokenString)
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, http2.BaseResponse[any]{
					Code:    http.StatusUnauthorized,
					Message: "Invalid or expired token",
				})
				c.Abort()
				return
			}
			fmt.Printf("9999----%+v\n", token.Claims.(jwt.MapClaims))
			userID := token.Claims.(jwt.MapClaims)["sub"].(string)
			c.Set("userID", userID) // 存储用户ID到上下文
			ctx := context.WithValue(c.Request.Context(), "userID", userID)
			c.Request = c.Request.WithContext(ctx)
		}

		c.Next()
	}
}
