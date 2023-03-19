package repository

import (
	mongoClient "cb-ldp-backend/commons/mongo"
	"cb-ldp-backend/config"
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"context"
	"log"

	_ "github.com/golang/mock/mockgen/model"
)

type MongoRepository struct {
	mongoClient mongoClient.MongoClient
}

type IMongoRepository interface {
	CreateQuizModule(ctx context.Context, module *entity.Module) error
	GetAllModules(ctx context.Context, isRequired bool) ([]entity.Module, error)
	GetModule(ctx context.Context, moduleId string, isRequired bool) (entity.Module, error)
	DeleteModule(ctx context.Context, moduleId interface{}) error
	FetchTestResult(ctx context.Context, moduleId string) ([]response.TestDetailsResponse, error)
	UpdateQuizModule(ctx context.Context, module entity.Module) (string, int, error)
	GetUserResult(ctx context.Context, moduleId string, email string) (interface{}, error)
	CreateQuiz(ctx context.Context, module *entity.Module) error
	DeleteQuiz(ctx context.Context, moduleId interface{}) error
	ViewQuiz(ctx context.Context, moduleId interface{}) ([]entity.Module, error)
	StoreResult(ctx context.Context, email string, quizResponse entity.QuizResponse) (interface{}, int, error)
	CreateQuestions(ctx context.Context, quiz *entity.Quiz) error
	DeleteQuestions(ctx context.Context, quizId interface{}) error
	DeleteQuestion(ctx context.Context, questionId interface{}) error
	GetCorrectAnswer(ctx context.Context, questionId interface{}, optionId interface{}) (entity.Option, string, error)
	CreateOptions(ctx context.Context, question entity.Question) error
	DeleteOptions(ctx context.Context, questionId interface{}) error
	CreateTokenOperation(ctx context.Context, user *entity.User) (error, string)
	GenerateToken(ctx context.Context, user entity.User) (error, string)
	GetUserRole(ctx context.Context, mail string) (error, string)
}

var envVar = config.LoadConfig()

func NewMongoRepository() IMongoRepository {
	mClient, err := mongoClient.ConnectDB()
	if err != nil {
		log.Fatal("Failed to initialize mongodb connection.", err.Error())
	}
	err = mClient.RunMigrations()
	if err != nil {
		log.Fatal("Failed to run migrations")
	}
	return &MongoRepository{
		mongoClient: mClient,
	}
}
