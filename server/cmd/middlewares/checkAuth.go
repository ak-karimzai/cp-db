package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ak-karimzai/cp-db/internal/jwt"
	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, ok := ctx.Request.Header["Authorization"]
		if !ok || len(tokenStr) == 0 {
			ctx.AbortWithError(
				http.StatusUnauthorized,
				errors.New("no authorization header"),
			)
			return
		}

		authHeaderContents := strings.Split(tokenStr[0], " ")
		if len(authHeaderContents) != 2 ||
			authHeaderContents[0] != "Bearer" ||
			authHeaderContents[1] == "" {
			ctx.AbortWithError(
				http.StatusUnauthorized,
				errors.New("invalid auth credintials"),
			)
			return
		}
		token := authHeaderContents[1]
		if _, err := jwt.ValidateToken(token); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				err,
			)
			return
		}
		ctx.Next()
	}
}
