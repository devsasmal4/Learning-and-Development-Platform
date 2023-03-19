package repository

import (
	"cb-ldp-backend/commons/utility"
	"cb-ldp-backend/models/entity"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m MongoRepository) CreateTokenOperation(ctx context.Context, user *entity.User) (error, string) {
	var token string
	var timeInDB int64 = user.TimeLoggedIn
	var timeAfterHours = timeInDB + 3600 // 12 hours - 43200
	filter := bson.M{"user_mail": user.UserMail}
	result := m.mongoClient.FindOne(ctx, "users", filter, nil)
	err := result.Decode(&user)
	if err == mongo.ErrNoDocuments {
		zohoId, empId, err := m.GetZohoEmpData(user.UserMail)
		if err != nil {
			log.Println("Error Verifying Data ->", err.Error())
			return err, token
		}
		user.Id = primitive.NewObjectID()
		user.EmployeeId = empId
		user.ZohoId = zohoId
		user.UserRole = "User"
		err = m.mongoClient.InsertOne(ctx, "users", user)
		err, token = m.GenerateToken(ctx, *user)
		if err != nil {
			log.Println("Error generating token", err.Error())
			return err, token
		}
	} else {
		if timeAfterHours > timeInDB {
			user.TimeLoggedIn = time.Now().Unix()
			filter := bson.M{"user_mail": user.UserMail}
			update := bson.M{"$set": bson.M{"time_logged_in": user.TimeLoggedIn}}
			err := m.mongoClient.FindOneAndUpdate(ctx, "users", filter, update).Decode(&user)
			if err != nil {
				log.Println("Error updating user", err.Error())
				return err, token
			}
			err, token = m.GenerateToken(ctx, *user)
		}
	}
	return nil, token
}

func (m MongoRepository) GenerateToken(ctx context.Context, user entity.User) (error, string) {
	filter := bson.M{"user_mail": user.UserMail}
	result := m.mongoClient.FindOne(ctx, "users", filter, nil)
	err := result.Decode(&user)
	message := user.UserMail + user.UserName + strconv.Itoa(int(user.TimeLoggedIn))
	key := []byte(envVar["secretKey"].(string))
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	token := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	if err == mongo.ErrNoDocuments {
		log.Println("Token not found", err.Error())
		return err, token
	}
	return nil, token
}

func (m MongoRepository) GetUserRole(ctx context.Context, mail string) (error, string) {
	var user entity.User
	filter := bson.M{"user_mail": mail}
	result := m.mongoClient.FindOne(ctx, "users", filter, nil)
	err := result.Decode(&user)
	if err == mongo.ErrNoDocuments {
		log.Println("Error decoding user", err.Error())
		return err, ""
	}
	return err, user.UserRole.String()
}

func (m MongoRepository) GetZohoEmpData(mail string) (int64, string, error) {

	refreshToken := envVar["refresh_token"]
	clientId := envVar["client_id"]
	clientSecret := envVar["client_secret"]
	url := fmt.Sprintf("https://accounts.zoho.com/oauth/v2/token?refresh_token=%v&client_id=%v&client_secret=%v&grant_type=refresh_token", refreshToken, clientId, clientSecret)
	resp, err := utility.HttpRequest("POST", url, nil, nil)
	if resp.StatusCode == 200 && err == nil {
		respBody, err := utility.ParseResponseBody(resp)
		accessTokenHeader := "Bearer " + respBody["access_token"]
		headers := map[string]string{"Authorization": strings.TrimSpace(accessTokenHeader)}
		url := fmt.Sprintf("https://people.zoho.com/api/forms/employee/getRecords?searchColumn=EMPLOYEEMAILALIAS&searchValue=%v", mail)
		resp, err := utility.HttpRequest("GET", url, nil, headers)
		if err != nil {
			log.Println("Error ->", err)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error Parsing the Zoho Response.")

		}
		bodyString := string(bodyBytes)
		zohoParams, err := utility.ProcessZohoResponse(bodyString)
		if err != nil {
			return 0, "", err
		}
		if zohoParams[0].EmailID != mail {
			return 0, "", errors.New("Email ID Mismatch.")
		}
		return zohoParams[0].ZohoID, zohoParams[0].EmployeeID, nil
	}
	return 0, "", err
}
