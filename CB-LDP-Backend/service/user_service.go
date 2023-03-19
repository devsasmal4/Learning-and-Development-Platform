package service

import (
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"context"
)

func (baseSvc *BaseService) GenerateToken(ctx context.Context, user entity.User) (interface{}, error) {
	err, token := baseSvc.mongoRepository.CreateTokenOperation(ctx, &user)
	if err != nil && token == "" {
		return nil, err
	}
	userResponse := response.UserResponse{
		Id:       user.Id,
		UserMail: user.UserMail,
		UserName: user.UserName,
		UserRole: user.UserRole,
		Token:    token,
	}
	return userResponse, nil
}
