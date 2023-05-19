package middlewares

import (
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/jwt"
	"github.com/gin-gonic/gin"
)

func BindUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("Authorization")
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}
		userInfo, responseErr := jwt.ValidateToken(token)
		if responseErr != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if err := ctx.ShouldBind(&userInfo); err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("user", userInfo)
		ctx.Next()
	}
}

func PageNotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{"error": "page not found"})
	}
}
