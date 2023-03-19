package handlers

import (
	"cb-ldp-backend/service"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	BaseSvc service.IBaseService
}

type IBaseHandler interface {
	CreateModule(*gin.Context)
	ViewAllModules(c *gin.Context)
	ViewModule(*gin.Context)
	ViewModuleInstructions(*gin.Context)
	UpdateModule(*gin.Context)
	DeleteModule(*gin.Context)
	GetTestDetails(*gin.Context)
	DownloadCsv(*gin.Context)
	GetUserResult(*gin.Context)
	// UpdateQuestion(*gin.Context)
	DeleteQuestion(*gin.Context)
	GetAnswer(*gin.Context)
	UploadCsv(*gin.Context)
	// CreateQuiz(*gin.Context)
	// UpdateQuiz(*gin.Context)
	// DeleteQuiz(*gin.Context)
	ViewQuiz(*gin.Context)
	ExecuteQuiz(*gin.Context)
	GenerateToken(*gin.Context)
}

func NewBaseHandler(baseService service.IBaseService) IBaseHandler {
	return &BaseHandler{
		BaseSvc: baseService,
	}
}
