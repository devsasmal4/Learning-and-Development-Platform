package service

import (
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/repository"

	"context"
)

type BaseService struct {
	mongoRepository repository.IMongoRepository
}

func NewBaseService(mongoRepository repository.IMongoRepository) IBaseService {
	return &BaseService{
		mongoRepository: mongoRepository,
	}
}

type IBaseService interface {
	CreateModule(context.Context, *entity.Module) error
	FetchAllModules(ctx context.Context) (interface{}, error)
	FetchModule(ctx context.Context, moduleId string) (interface{}, error)
	FetchModuleInstruction(ctx context.Context, moduleId string) (interface{}, error)
	UpdateModule(ctx context.Context, module entity.Module, moduleId string) (string, int, error)
	DeleteModule(ctx context.Context, moduleId string) (int, string, error)
	FetchTestDetails(ctx context.Context, moduleId string) (interface{}, error)
	GenerateCsv(ctx context.Context, moduleId string) (string, error)
	FetchUserResult(ctx context.Context, moduleId string, email string) (interface{}, error)
	// CreateQuiz(ctx context.Context) error
	// UpdateQuiz(ctx context.Context, quizId string) error
	// DeleteQuiz(ctx context.Context, quizId string) error
	ViewQuiz(ctx context.Context, moduleId string) (interface{}, error)
	CreateQuizResponse(ctx context.Context, email string, quizResponse entity.QuizResponse) (interface{}, int, error)
	// UpdateQuestion(ctx context.Context, questionId string, question entity.Question) error
	DeleteQuestion(ctx context.Context, questionId string) (string, int, error)
	FetchCorrectAnswer(ctx context.Context, questionId string, optionId string) (interface{}, error)
	GenerateToken(ctx context.Context, user entity.User) (interface{}, error)
}
