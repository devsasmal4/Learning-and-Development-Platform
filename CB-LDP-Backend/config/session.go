package config

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionStore(c *gin.Context) map[string]interface{} {
	session := sessions.Default(c)
	sessionEmail := session.Get("email")
	sessionUserName := session.Get("name")
	if sessionEmail == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
	}
	var result map[string]interface{}
	result["email"] = sessionEmail
	result["name"] = sessionUserName
	return result
}
