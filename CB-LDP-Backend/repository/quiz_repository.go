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

func (m MongoRepository) CreateQuiz(ctx context.Context, module *entity.Module) error {
	quiz := module.ModuleQuiz
	quiz.Id = primitive.NewObjectID()
	quiz.ModuleId = module.Id
	err := m.CreateQuestions(ctx, &quiz)
	if err != nil {
		log.Println("Error creating questions", err.Error())
		return err

	}

	err = m.mongoClient.InsertOne(ctx, "quizzes", quiz)
	if err != nil {
		log.Println("Error inserting", err.Error())
		return err
	}
	module.ModuleQuiz = quiz
	return nil
}

func (m MongoRepository) DeleteQuiz(ctx context.Context, moduleId interface{}) error {
	opts := options.Find().SetProjection(bson.D{{"_id", 1}})
	cursor, err := m.mongoClient.Find(ctx, "quizzes", bson.M{"module_id": moduleId}, opts)
	if err != nil {
		log.Println("Quiz not found", err.Error())
		return err
	}
	var quizIds []bson.D
	cursor.All(ctx, &quizIds)
	err = m.mongoClient.DeleteMany(ctx, "quizzes", bson.M{"module_id": moduleId})
	if err != nil {
		log.Println("Error deleting quizes", err.Error())
		return err
	}
	for _, quizId := range quizIds {
		id := quizId[0].Value
		err = m.DeleteQuestions(ctx, id)
		if err != nil {
			log.Println("Error deleting questions", err.Error())
			return err
		}
	}
	return nil
}

func (m MongoRepository) ViewQuiz(ctx context.Context, moduleId interface{}) ([]entity.Module, error) {
	opts := options.Find().SetProjection(bson.D{{"module_quiz.quiz_questionaires.answer_explanation", 0}, {"module_quiz.quiz_questionaires.question_options.option_score", 0}, {"module_quiz.quiz_questionaires.question_options.is_correct", 0}})
	cursor, err := m.mongoClient.Find(ctx, "modules", bson.M{"_id": moduleId}, opts)
	if err != nil {
		log.Println("Module not found", err.Error())
		return []entity.Module{}, err
	}
	var modules []entity.Module
	cursor.All(ctx, &modules)
	if modules == nil {
		return []entity.Module{}, errors.New("Invalid module Id")
	}
	return modules, nil
}

func (m MongoRepository) StoreResult(ctx context.Context, email string, quizResponse entity.QuizResponse) (interface{}, int, error) {
	var updatedQuizResponse entity.QuizResponse
	var user entity.User
	if err := m.mongoClient.FindOne(ctx, "users", bson.M{"user_mail": email}, nil).Decode(&user); err != nil {

		log.Println("User not found", err.Error())
		return "User not found", 404, err
	}
	quizResponse.UserId = user.Id

	var module entity.Module
	if err := m.mongoClient.FindOne(ctx, "modules", bson.M{"_id": quizResponse.ModuleId}, nil).Decode(&module); err != nil {
		log.Println("Module not found", err.Error())
		return "Module not found", 404, err
	}

	filter := bson.M{"user_id": user.Id, "module_id": module.Id}
	if err := m.mongoClient.FindOne(ctx, "quiz_responses", filter, nil).Decode(&updatedQuizResponse); err == mongo.ErrNoDocuments {
		quizResponse.Id = primitive.NewObjectID()
		err := m.mongoClient.InsertOne(ctx, "quiz_responses", quizResponse)
		if err != nil {
			log.Println("Error creating quiz response", err.Error())
			return "Error creating quiz response", 500, err
		}
	} else {
		filters := bson.M{"user_id": user.Id, "module_id": module.Id}
		update := bson.M{"$set": bson.M{"quiz_score": quizResponse.QuizScore, "quiz_result": quizResponse.QuizResult, "start_date": quizResponse.StartDate, "end_date": quizResponse.EndDate}}
		err := m.mongoClient.FindOneAndUpdate(ctx, "quiz_responses", filters, update).Decode(&updatedQuizResponse)
		if err != nil {
			log.Println("Error updating quiz responses", err.Error())
			return "Error updating quiz response", 404, err
		}
	}
	quizResponse.Id = updatedQuizResponse.Id

	return quizResponse, 200, nil
}
