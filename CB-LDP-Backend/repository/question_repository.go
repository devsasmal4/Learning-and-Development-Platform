package repository

import (
	"cb-ldp-backend/models/entity"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongoRepository) CreateQuestions(ctx context.Context, quiz *entity.Quiz) error {
	questions := quiz.QuizQuestionaires
	for index, question := range questions {
		question.Id = primitive.NewObjectID()
		question.QuizId = quiz.Id
		quiz.QuizQuestionaires[index] = question
		err := m.CreateOptions(ctx, question)
		if err != nil {
			log.Println("Error creating options", err.Error())
			return err
		}

		err = m.mongoClient.InsertOne(ctx, "questions", question)
		if err != nil {
			log.Println("Error inserting questions", err.Error())
			return err
		}
	}
	return nil
}

func (m MongoRepository) DeleteQuestions(ctx context.Context, quizId interface{}) error {
	opts := options.Find().SetProjection(bson.D{{"_id", 1}})
	cursor, err := m.mongoClient.Find(ctx, "questions", bson.M{"quiz_id": quizId}, opts)
	if err != nil {
		log.Println("Error finding questions", err.Error())
		return err
	}
	var questionIds []bson.D
	cursor.All(ctx, &questionIds)
	err = m.mongoClient.DeleteMany(ctx, "questions", bson.M{"quiz_id": quizId})
	if err != nil {
		log.Println("Error deleting questions", err.Error())
		return err
	}

	for _, questionId := range questionIds {
		id := questionId[0].Value
		err = m.DeleteOptions(ctx, id)
		if err != nil {
			log.Println("Error deleting options", err.Error())
			return err
		}
	}
	return nil
}

func (m MongoRepository) DeleteQuestion(ctx context.Context, questionId interface{}) error {
	filter := bson.M{"_id": questionId}
	err := m.mongoClient.FindOneAndDelete(ctx, "questions", filter)
	if err != nil {
		log.Println("Delete Error", err.Error())
		return err
	}
	return nil
}

func (m MongoRepository) GetCorrectAnswer(ctx context.Context, questionId interface{}, optionId interface{}) (entity.Option, string, error) {
	var question entity.Question
	var option entity.Option
	if err := m.mongoClient.FindOne(ctx, "questions", bson.M{"_id": questionId}, nil).Decode(&question); err != nil && err == mongo.ErrNoDocuments {
		log.Println("Question not found", err.Error())
		return entity.Option{}, "Question not found", err
	}

	if err := m.mongoClient.FindOne(ctx, "options", bson.M{"_id": optionId}, nil).Decode(&option); err != nil && err == mongo.ErrNoDocuments {
		log.Println("Option not found", err.Error())
		return entity.Option{}, "Option not found", err
	}
	var correctOption entity.Option
	if questionId != option.QuestionId {
		return entity.Option{}, "Option belongs to different question", errors.New("Invalid option")
	}
	if !option.IsCorrect {
		if err := m.mongoClient.FindOne(ctx, "options", bson.M{"question_id": option.QuestionId, "is_correct": true}, nil).Decode(&correctOption); err != nil && err == mongo.ErrNoDocuments {
			return entity.Option{}, "No correct answer", errors.New("Invalid option")
		}
		option.OptionString = correctOption.OptionString
	}
	return option, question.AnswerExplanation, nil
}
