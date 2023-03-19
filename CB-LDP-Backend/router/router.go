package router

import (
	"cb-ldp-backend/config"
	_ "cb-ldp-backend/docs"
	"cb-ldp-backend/handlers"
	"cb-ldp-backend/middleware"
	serverError "cb-ldp-backend/models/error"
	"cb-ldp-backend/repository"
	"cb-ldp-backend/service"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Coffbeans Quiz API
// @version 1.0
// @description Coffeebeans learning and development.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
// @schemes http
var envVar = config.LoadConfig()

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.EnableCors())

	mongoRepository := repository.NewMongoRepository()
	baseService := service.NewBaseService(mongoRepository)
	baseHandler := handlers.NewBaseHandler(baseService)

	adminRouter := router.Group("/", middleware.TokenAuthentication(mongoRepository), middleware.CheckRole(mongoRepository))
	module := adminRouter.Group("/module")
	{
		module.POST("", baseHandler.CreateModule)
		module.PUT("/:moduleId", baseHandler.UpdateModule)
		module.DELETE("/:moduleId", baseHandler.DeleteModule)
		module.GET("/:moduleId/testDetails", baseHandler.GetTestDetails)
		module.GET("/:moduleId/download", baseHandler.DownloadCsv)
	}

	quiz := adminRouter.Group("/quiz")
	{
		// quiz.POST("/", baseHandler.CreateQuiz)
		// quiz.PUT("/:quizId", baseHandler.UpdateQuiz)
		//quiz.DELETE("/:quizId", baseHandler.DeleteQuiz)
	}

	question := adminRouter.Group("/question")
	{
		// question.PUT("/:questionId", baseHandler.UpdateQuestion)
		question.DELETE("/:questionId", baseHandler.DeleteQuestion)
		question.POST("/upload", baseHandler.UploadCsv)
	}

	userRouter := router.Group("/", middleware.TokenAuthentication(mongoRepository))
	module = userRouter.Group("/module")
	{
		module.GET("", baseHandler.ViewAllModules)
		module.GET("/:moduleId/instructions", baseHandler.ViewModuleInstructions)
		module.GET("/:moduleId", baseHandler.ViewModule)
		module.GET("/:moduleId/userResult", baseHandler.GetUserResult)
	}

	quiz = userRouter.Group("/quiz")
	{
		quiz.GET("/:moduleId", baseHandler.ViewQuiz)
		quiz.POST("/execute", baseHandler.ExecuteQuiz)
	}

	question = userRouter.Group("/question")
	{
		question.GET("/:questionId/answer/:optionId", baseHandler.GetAnswer)
	}

	user := router.Group("/user")
	{
		user.POST("/createToken", baseHandler.GenerateToken)
	}

	url := ginSwagger.URL(fmt.Sprintf("http://%v/swagger/doc.json", envVar["host"]))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, serverError.HandleNoRouteError())
	})

	return router

}
