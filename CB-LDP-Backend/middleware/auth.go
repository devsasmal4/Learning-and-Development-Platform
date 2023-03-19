package middleware

import (
	"cb-ldp-backend/constants"
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/repository"

	error2 "cb-ldp-backend/models/error"
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenAuthentication(m repository.IMongoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		var token = c.GetHeader("Token")
		user.UserMail = c.GetHeader("Email")
		if token == "" || user.UserMail == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, error2.HandleUnauthorizedError(errors.New("Invalid headers")))
			return
		}

		if strings.Contains(user.UserMail, "@") && strings.Split(user.UserMail, "@")[1] != "coffeebeans.io" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, error2.HandleUnauthorizedError(errors.New("Invalid email id / not a coffeebeans email id")))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
		defer cancel()
		err, generatedToken := m.GenerateToken(ctx, user)
		if err != nil || generatedToken != token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, error2.HandleUnauthorizedError(errors.New("Invalid token")))
			return
		}
		c.Next()
	}
}

func CheckRole(m repository.IMongoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user entity.User
		user.UserMail = c.GetHeader("Email")
		ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
		defer cancel()
		err, role := m.GetUserRole(ctx, user.UserMail)
		if role == "Admin" || role == "Module Owner" && err == nil {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, error2.HandleUnauthorizedError(errors.New("User has no access")))
		}
	}
}
