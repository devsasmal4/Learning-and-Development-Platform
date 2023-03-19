package service

import (
	"cb-ldp-backend/commons/utility"
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var quiz entity.Quiz

// func (baseSvc *BaseService) CreateQuiz(ctx context.Context) error {
// 	quiz.Id = primitive.NewObjectID()
// 	err := quiz.CreateQuestions(ctx, baseSvc.mongoRepository)
// 	if err != nil {
// 		return err
// 	}

// 	err = baseSvc.mongoRepository.InsertOne(ctx, "quizzes", quiz)
// 	if err != nil {
// 		return nil
// 	}
// 	return nil
// }

// func (baseSvc *BaseService) UpdateQuiz(ctx context.Context, quizId string) error {
// 	objId, _ := primitive.ObjectIDFromHex(quizId)
// 	quiz.Id = objId
// 	update := bson.M{"$set": quiz}
// 	filter := bson.M{"_id": objId}
// 	var updatedQuiz entity.Quiz
// 	err := baseSvc.mongoRepository.FindOneAndUpdate(ctx, "quizzes", filter, update).Decode(&updatedQuiz)
// 	for _, question := range quiz.QuizQuestionaires {
// 		baseSvc.mongoRepository.FindOneAndUpdate(ctx, "questions", bson.M{"_id": question.Id}, bson.M{"$set": question})
// 		for _, option := range question.QuestionOptions {
// 			baseSvc.mongoRepository.FindOneAndUpdate(ctx, "options", bson.M{"_id": option.Id}, bson.M{"$set": option})
// 		}
// 	}
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (baseSvc *BaseService) DeleteQuiz(ctx context.Context, quizId string) error {
// 	objId, _ := primitive.ObjectIDFromHex(quizId)
// 	filter := bson.M{"_id": objId}
// 	err := baseSvc.mongoRepository.FindOneAndDelete(ctx, "quizzes", filter)
// 	if err != nil {
// 		return nil
// 	}

// 	err = commons.DeleteQuestionsAndOptions(ctx, baseSvc.mongoRepository, objId)
// 	return nil
// }

func (baseSvc *BaseService) ViewQuiz(ctx context.Context, moduleId string) (interface{}, error) {
	objId, _ := primitive.ObjectIDFromHex(moduleId)

	module, err := baseSvc.mongoRepository.ViewQuiz(ctx, objId)
	if err != nil {
		return nil, err
	}

	quizQuestions := response.QuizQuestions{
		ModuleName:        module[0].ModuleName,
		QuizQuestionaires: module[0].ModuleQuiz.QuizQuestionaires,
	}
	utility.Shuffle(quizQuestions)

	return quizQuestions, nil
}

func (baseSvc *BaseService) CreateQuizResponse(ctx context.Context, email string, quizResponse entity.QuizResponse) (interface{}, int, error) {
	data, code, err := baseSvc.mongoRepository.StoreResult(ctx, email, quizResponse)
	if err != nil {
		return data, code, err
	}
	return data, 200, nil
}
