package repository

import (
	"cb-ldp-backend/commons/utility"
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "modules"

func (m MongoRepository) CreateQuizModule(ctx context.Context, module *entity.Module) error {
	err := m.mongoClient.InsertOne(ctx, "modules", &module)
	if err != nil {
		log.Println("Error inserting module", err.Error())
		return err
	}
	return nil
}

func (m MongoRepository) GetAllModules(ctx context.Context, isRequired bool) ([]entity.Module, error) {
	var modules []entity.Module
	opts := options.Find().SetProjection(bson.M{"_id": 1, "module_name": 1, "module_created_on": 1})
	results, err := m.mongoClient.Find(ctx, collectionName, bson.M{}, opts)
	if err != nil {
		log.Println("Module not found", err.Error())
		return nil, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleModule entity.Module
		if err := results.Decode(&singleModule); err != nil {
			log.Println("Error decoding modules", err.Error())
			return nil, err
		}
		modules = append(modules, singleModule)
	}
	return modules, nil
}

func (m MongoRepository) GetModule(ctx context.Context, moduleId string, isRequired bool) (entity.Module, error) {
	objId, _ := primitive.ObjectIDFromHex(moduleId)
	filter := bson.M{"_id": objId}
	var module entity.Module
	opts := options.FindOne().SetProjection(bson.M{"_id": 1, "module_name": 1, "module_study_material": !isRequired, "module_description": !isRequired, "module_instructions": isRequired})
	results := m.mongoClient.FindOne(ctx, collectionName, filter, opts)

	if err := results.Decode(&module); err != nil {
		log.Println("Error decoding module", err.Error())
		return entity.Module{}, err
	}
	return module, nil
}

func (m MongoRepository) DeleteModule(ctx context.Context, moduleId interface{}) error {
	filter := bson.M{"_id": moduleId}
	err := m.mongoClient.FindOneAndDelete(ctx, "modules", filter)
	if err != nil {
		log.Println("Module not found", err.Error())
		return err
	}
	return nil
}

func (m MongoRepository) FetchTestResult(ctx context.Context, moduleId string) ([]response.TestDetailsResponse, error) {
	var testDetails []response.TestDetailsResponse
	var module entity.Module
	objId, _ := primitive.ObjectIDFromHex(moduleId)

	opts := options.FindOne().SetProjection(bson.M{"_id": 1, "module_name": 1})
	result := m.mongoClient.FindOne(ctx, "modules", bson.M{"_id": objId}, opts)
	if err := result.Decode(&module); err != nil {
		log.Println("Error decoding module", err.Error())
		return nil, err
	}
	results_responses, err := m.mongoClient.Find(ctx, "quiz_responses", bson.M{"module_id": objId}, nil)
	if err != nil {
		log.Println("Module not found", err.Error())
		return nil, err
	}
	defer results_responses.Close(ctx)

	for results_responses.Next(ctx) {
		var quizResponse entity.QuizResponse
		if err := results_responses.Decode(&quizResponse); err != nil {
			log.Println("Error decoding quizresponse", err.Error())
			return nil, err
		}
		var user entity.User
		opts := options.FindOne().SetProjection(bson.M{"_id": 1, "user_name": 1})
		result_user := m.mongoClient.FindOne(ctx, "users", bson.M{"_id": quizResponse.UserId}, opts)
		if err := result_user.Decode(&user); err != nil {
			log.Println("Error decoding user", err.Error())
			return nil, err
		}

		duration, testStatus := utility.GetDurationAndStatus(quizResponse.StartDate, quizResponse.EndDate)

		testDetailResponse := response.TestDetailsResponse{
			Id:         quizResponse.Id,
			UserName:   user.UserName,
			ModuleName: module.ModuleName,
			TestStatus: testStatus,
			QuizScore:  quizResponse.QuizScore,
			QuizResult: quizResponse.QuizResult,
			StartDate:  quizResponse.StartDate,
			EndDate:    quizResponse.EndDate,
			Duration:   duration,
		}
		testDetails = append(testDetails, testDetailResponse)
	}
	return testDetails, nil
}

func (m MongoRepository) UpdateQuizModule(ctx context.Context, module entity.Module) (string, int, error) {
	var updatedModule entity.Module
	var updatedQuiz entity.Quiz
	var updatatedQuestion entity.Question
	var updatedOption entity.Option

	filter := bson.M{"_id": module.Id}
	update := bson.M{"$set": module}
	if err := m.mongoClient.FindOne(ctx, "modules", bson.M{"_id": module.Id}, nil).Decode(&updatedModule); err != nil {
		log.Println("Module not found", err.Error())
		return "Module not found", 404, err
	}

	err := m.mongoClient.FindOneAndUpdate(ctx, "modules", filter, update).Decode(&updatedModule)
	if err != nil {
		log.Println("Error updating module", err.Error())
		return "Error updating module", 500, err
	}

	err = m.mongoClient.FindOneAndUpdate(ctx, "quizzes", bson.M{"_id": module.ModuleQuiz.Id}, bson.M{"$set": module.ModuleQuiz}).Decode(&updatedQuiz)
	if err != nil {
		log.Println("Error updating quiz", err.Error())
		return "Error updating quiz", 500, err
	}

	for _, question := range module.ModuleQuiz.QuizQuestionaires {
		err = m.mongoClient.FindOneAndUpdate(ctx, "questions", bson.M{"_id": question.Id}, bson.M{"$set": question}).Decode(&updatatedQuestion)
		if err != nil {
			log.Println("Error updating questions", err.Error())
			return "Error updating questions", 500, err
		}

		for _, option := range question.QuestionOptions {
			err = m.mongoClient.FindOneAndUpdate(ctx, "options", bson.M{"_id": option.Id}, bson.M{"$set": option}).Decode(&updatedOption)
			if err != nil {
				log.Println("Error updating options", err.Error())
				return "Error updating options", 500, err
			}
		}
	}
	return "Successfully updated module", 200, nil
}

func (m MongoRepository) GetUserResult(ctx context.Context, moduleId string, email string) (interface{}, error) {
	var module entity.Module

	objId, _ := primitive.ObjectIDFromHex(moduleId)

	opts := options.FindOne().SetProjection(bson.M{"_id": 1, "module_name": 1})
	if err := m.mongoClient.FindOne(ctx, "modules", bson.M{"_id": objId}, opts).Decode(&module); err != nil {
		log.Println("Module not found", err.Error())
		return "Module not found", err
	}

	var user entity.User
	if err := m.mongoClient.FindOne(ctx, "users", bson.M{"user_mail": email}, nil).Decode(&user); err != nil {
		log.Println("User not found", err.Error())
		return "User not found", err
	}

	filter := bson.M{"user_id": user.Id, "module_id": module.Id}
	var quizResponse entity.QuizResponse
	if err := m.mongoClient.FindOne(ctx, "quiz_responses", filter, nil).Decode(&quizResponse); err != nil {
		return "User result not found", err
	}

	duration, testStatus := utility.GetDurationAndStatus(quizResponse.StartDate, quizResponse.EndDate)
	testDetailResponse := response.TestDetailsResponse{
		Id:         quizResponse.Id,
		UserName:   user.UserName,
		ModuleName: module.ModuleName,
		TestStatus: testStatus,
		QuizScore:  quizResponse.QuizScore,
		QuizResult: quizResponse.QuizResult,
		StartDate:  quizResponse.StartDate,
		EndDate:    quizResponse.EndDate,
		Duration:   duration,
	}
	return testDetailResponse, nil
}
