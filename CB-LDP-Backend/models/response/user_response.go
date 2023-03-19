package response

import (
	"cb-ldp-backend/constants"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	Id       primitive.ObjectID `bson:"_id"`
	UserMail string             `bson:"user_mail" json:"user_mail"`
	UserName string             `bson:"user_name" json:"user_name"`
	UserRole constants.UserRole `bson:"user_role" json:"user_role"`
	Token    string             `bson:"token" json:"token"`
}
