package middleware

import (
	"cb-ldp-backend/commons/utility"
	"cb-ldp-backend/constants"
	error2 "cb-ldp-backend/models/error"
	"errors"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

func EnableCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get(constants.Origin)
		allowedOrigins := constants.AllowedOrigin
		if len(origin) != 0 && (utility.Contains(allowedOrigins, origin)) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			ctx.Writer.Header().Set("Access-Control-Allow-Headers",
				strings.Join([]string{
					constants.Origin,
					constants.Accept,
					constants.ContentTypeHeader,
					constants.Authorization,
					constants.DateUsed,
					constants.XRequestedWith,
					constants.Cookie,
					constants.Email,
					constants.Token,
				}, ","))
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				error2.HandleUnauthorizedError(errors.New("origin not allowed")))
		}

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
		}
		ctx.Next()
	}
}
