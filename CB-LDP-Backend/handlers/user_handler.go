package handlers

import (
	"cb-ldp-backend/constants"
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

var uCollectionName string = "users"

// @Summary      Create token
// @Description  Create token
// @Produce      json
// @Param entity.User  body entity.User true  "User request body"
// @Success      200  {object}  response.UserResponse
// @failure      400              {string}    "Bad Request"
// @failure      500              {string}    "Server error"
// @Router       /user/createToken [post]
func (h *BaseHandler) GenerateToken(c *gin.Context) {
	var user entity.User
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	user.TimeLoggedIn = time.Now().Unix()
	if err := c.ShouldBindJSON(&user); err != nil {
		response.APIResponse(c, err.Error(), 400, "Bad Request")
		return
	}

	userResponse, err := h.BaseSvc.GenerateToken(ctx, user)
	if err != nil {
		response.APIResponse(c, err.Error(), 500, "Server Error")
		return
	}

	response.APIResponse(c, userResponse, 200, "User Token Created")
}
