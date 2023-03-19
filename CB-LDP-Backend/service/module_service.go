package service

import (
	"cb-ldp-backend/commons/utility"
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionName string = "modules"

func (baseSvc *BaseService) CreateModule(ctx context.Context, module *entity.Module) error {
	module.Id = primitive.NewObjectID()
	err := baseSvc.mongoRepository.CreateQuiz(ctx, module)
	if err != nil {
		return err
	}

	err = baseSvc.mongoRepository.CreateQuizModule(ctx, module)
	if err != nil {
		return err
	}
	return nil
}

func (baseSvc *BaseService) FetchAllModules(ctx context.Context) (interface{}, error) {
	modules, err := baseSvc.mongoRepository.GetAllModules(ctx, false)
	if err != nil {
		return nil, err
	}
	var allModuleResponses []response.ModuleResponse
	for index := range modules {
		singleModuleResponse := response.ModuleResponse{
			Id:              modules[index].Id,
			ModuleName:      modules[index].ModuleName,
			ModuleCreatedOn: modules[index].ModuleCreatedOn,
		}
		allModuleResponses = append(allModuleResponses, singleModuleResponse)
	}
	return allModuleResponses, nil
}

func (baseSvc *BaseService) FetchModule(ctx context.Context, moduleId string) (interface{}, error) {
	module, err := baseSvc.mongoRepository.GetModule(ctx, moduleId, false)
	if err != nil {
		return nil, err
	}
	moduleResponse := response.ModuleViewResponse{
		Id:                  module.Id,
		ModuleName:          module.ModuleName,
		ModuleStudyMaterial: module.ModuleStudyMaterial,
		ModuleDescription:   module.ModuleDescription,
	}
	return moduleResponse, nil
}

func (baseSvc *BaseService) FetchModuleInstruction(ctx context.Context, moduleId string) (interface{}, error) {
	module, err := baseSvc.mongoRepository.GetModule(ctx, moduleId, true)
	if err != nil {
		return nil, err
	}
	instructionResponse := response.ModuleInstructionsResponse{
		Id:                 module.Id,
		ModuleName:         module.ModuleName,
		ModuleInstructions: module.ModuleInstructions,
	}
	return instructionResponse, nil
}

func (baseSvc *BaseService) UpdateModule(ctx context.Context, module entity.Module, moduleId string) (string, int, error) {
	objId, _ := primitive.ObjectIDFromHex(moduleId)
	module.Id = objId
	message, code, err := baseSvc.mongoRepository.UpdateQuizModule(ctx, module)
	return message, code, err
}

func (baseSvc *BaseService) DeleteModule(ctx context.Context, moduleId string) (int, string, error) {
	objId, _ := primitive.ObjectIDFromHex(moduleId)
	err := baseSvc.mongoRepository.DeleteModule(ctx, objId)
	if err != nil {
		return 404, "Module not found", err
	}
	err = baseSvc.mongoRepository.DeleteQuiz(ctx, objId)
	if err != nil {
		return 500, "Error deleting quiz, questions and options", err
	}
	return 200, "Sucessfully deleted Module", nil
}

func (baseSvc *BaseService) FetchTestDetails(ctx context.Context, moduleId string) (interface{}, error) {
	testDetails, err := baseSvc.mongoRepository.FetchTestResult(ctx, moduleId)
	if err != nil {
		return nil, err
	}
	return testDetails, nil
}

func (baseSvc *BaseService) GenerateCsv(ctx context.Context, moduleId string) (string, error) {
	testDetails, err := baseSvc.mongoRepository.FetchTestResult(ctx, moduleId)
	if err != nil {
		return "Error fetching data", err
	}
	path, err := utility.GenerateCsv(testDetails)
	if err != nil {
		return "Error generating data", err
	}
	return path, nil
}

func (baseSvc *BaseService) FetchUserResult(ctx context.Context, moduleId string, email string) (interface{}, error) {
	result, err := baseSvc.mongoRepository.GetUserResult(ctx, moduleId, email)
	return result, err

}
