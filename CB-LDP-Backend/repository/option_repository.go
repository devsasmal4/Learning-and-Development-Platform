package repository

import (
	"cb-ldp-backend/models/entity"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongoRepository) CreateOptions(ctx context.Context, question entity.Question) error {
	options := question.QuestionOptions
	for index, option := range options {
		option.Id = primitive.NewObjectID()
		option.QuestionId = question.Id
		err := m.mongoClient.InsertOne(ctx, "options", option)
		if err != nil {
			log.Println("Error inserting option", err.Error())
			return err
		}
		question.QuestionOptions[index] = option
	}
	return nil
}

func (m MongoRepository) DeleteOptions(ctx context.Context, questionId interface{}) error {
	opts := options.Find().SetProjection(bson.D{{"_id", 1}})
	cursor, err := m.mongoClient.Find(ctx, "options", bson.M{"question_id": questionId}, opts)
	if err != nil {
		log.Println("Error finding options", err.Error())
		return err
	}
	var questionIds []bson.D
	cursor.All(ctx, &questionIds)
	err = m.mongoClient.DeleteMany(ctx, "options", bson.M{"question_id": questionId})
	if err != nil {
		log.Println("Error deleting options", err.Error())
		return err
	}
	return nil
}
