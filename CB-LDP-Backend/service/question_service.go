package service

import (
	"cb-ldp-backend/models/response"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func (baseSvc *BaseService) UpdateQuestion(ctx context.Context, questionId string, question entity.Question) error {
// 	objId, _ := primitive.ObjectIDFromHex(questionId)

// 	question.Id = objId
// 	update := bson.M{"$set": question}
// 	filter := bson.M{"_id": questionId}
// 	var updatedQuestion entity.Question
// 	err := baseSvc.mongoRepository.FindOneAndUpdate(ctx, "questions", filter, update).Decode(&updatedQuestion)
// 	for _, option := range question.QuestionOptions {
// 		option.QuestionId = objId
// 		baseSvc.mongoRepository.FindOneAndUpdate(ctx, "options", bson.M{"_id": option.Id}, bson.M{"$set": option})
// 	}
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (baseSvc *BaseService) DeleteQuestion(ctx context.Context, questionId string) (string, int, error) {
	objId, _ := primitive.ObjectIDFromHex(questionId)

	err := baseSvc.mongoRepository.DeleteQuestion(ctx, objId)
	if err != nil {
		return "Question not found", 404, err
	}

	err = baseSvc.mongoRepository.DeleteOptions(ctx, objId)
	if err != nil {
		return "Error deleting options", 500, err
	}
	return "Successfully deleted question", 200, nil
}

func (baseSvc *BaseService) FetchCorrectAnswer(ctx context.Context, questionId string, optionId string) (interface{}, error) {
	questionObjId, _ := primitive.ObjectIDFromHex(questionId)
	optionObjId, _ := primitive.ObjectIDFromHex(optionId)

	option, data, err := baseSvc.mongoRepository.GetCorrectAnswer(ctx, questionObjId, optionObjId)
	if err != nil {
		return data, err
	}
	answerResponse := response.AnswerResponse{
		ScoreEarned:       option.OptionScore,
		AnswerExplanation: data,
		CorrectAnswer:     option.OptionString,
	}
	return answerResponse, err
}
