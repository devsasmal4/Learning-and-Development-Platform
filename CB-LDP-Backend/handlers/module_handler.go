package handlers

import (
	"cb-ldp-backend/constants"

	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

var collectionName string = "modules"

// @Summary      Create module
// @Description  Create module
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param entity.Module  body entity.Module true  "Module request body"
// @Success      200  {object}  entity.Module
// @failure      400              {string}    "Bad Request"
// @failure      500              {string}    "Error creating quiz"
// @Router       /module [post]
func (h *BaseHandler) CreateModule(c *gin.Context) {
	var module entity.Module
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&module); err != nil {
		response.APIResponse(c, err.Error(), 400, "Error")
		return
	}

	err := h.BaseSvc.CreateModule(ctx, &module)
	if err != nil {
		response.APIResponse(c, err.Error(), 500, "Error creating quiz")
		return
	}

	response.APIResponse(c, nil, 200, "Module Created")
}

// @Summary      View all modules
// @Description  View all modules
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Success      200  {object}  []response.ModuleViewResponse
// @failure      404             {string}    "Modules not found"
// @failure      500              {string}    "Server Error"
// @Router       /module [get]
func (h *BaseHandler) ViewAllModules(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	allModuleResponses, err := h.BaseSvc.FetchAllModules(ctx)
	if err != nil {
		response.APIResponse(c, err.Error(), 404, "Modules not found")
		return
	}
	response.APIResponse(c, allModuleResponses, 200, "Module List displayed")
}

// @Summary      View module
// @Description  View module
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param moduleId path string true "Module Id"
// @Success      200  {object}  response.ModuleViewResponse
// @failure      404              {string}    "Module not found"
// @Router       /module/{moduleId} [get]
func (h *BaseHandler) ViewModule(c *gin.Context) {
	moduleId := c.Param("moduleId")
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	moduleResponse, err := h.BaseSvc.FetchModule(ctx, moduleId)
	if err != nil {
		response.APIResponse(c, err.Error(), 404, "Module Not Found")
		return
	}
	response.APIResponse(c, moduleResponse, 200, "Module displayed")
}

// @Summary      View module instructions
// @Description  View module instructions
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param moduleId path string true "Module Id"
// @Success      200  {object}  response.ModuleInstructionsResponse
// @failure      404              {string}    "Module not found"
// @Router       /module/{moduleId}/instructions [get]
func (h *BaseHandler) ViewModuleInstructions(c *gin.Context) {
	moduleId := c.Param("moduleId")
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()

	instructionResponse, err := h.BaseSvc.FetchModuleInstruction(ctx, moduleId)
	if err != nil {
		response.APIResponse(c, err.Error(), 404, "Module not found")
		return
	}

	response.APIResponse(c, instructionResponse, 200, "Module List displayed")
}

// @Summary      Update module
// @Description  Update module
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param moduleId path string true "Module Id"
// @Param entity.Module  body entity.Module true  "Updated module request body"
// @Success      200  {object}  response.Response
// @failure      400              {string}    "Bad Request"
// @failure      404              {string}    "Module not found"
// @failure      500              {string}    "Error updating module"
// @Router       /module/{moduleId} [put]
func (h *BaseHandler) UpdateModule(c *gin.Context) {
	var module entity.Module
	moduleId := c.Param("moduleId")
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&module); err != nil {
		response.APIResponse(c, err.Error(), 400, "Error")
		return
	}

	message, code, err := h.BaseSvc.UpdateModule(ctx, module, moduleId)
	if err != nil {
		response.APIResponse(c, err.Error(), code, message)
		return
	}
	response.APIResponse(c, nil, code, message)
}

func (h *BaseHandler) DeleteModule(c *gin.Context) {
	moduleId := c.Param("moduleId")
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	statusCode, message, err := h.BaseSvc.DeleteModule(ctx, moduleId)
	if err != nil {
		response.APIResponse(c, err.Error(), statusCode, message)
		return
	}
	response.APIResponse(c, nil, statusCode, message)
}

// @Summary      Get test details
// @Description  Get test details
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param moduleId path string true "Module Id"
// @Success      200  {object}  []response.TestDetailsResponse
// @failure      404              {string}    "Module not found"
// @failure      500              {string}    "Server Error"
// @Router       /module/{moduleId}/testDetails [get]
func (h *BaseHandler) GetTestDetails(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	testDetails, err := h.BaseSvc.FetchTestDetails(ctx, c.Param("moduleId"))

	if err != nil {
		response.APIResponse(c, err.Error(), 404, "Record not found")
		return
	}

	response.APIResponse(c, testDetails, 200, "Details Displayed")
}

func (h *BaseHandler) DownloadCsv(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()

	data, err := h.BaseSvc.GenerateCsv(ctx, c.Param("moduleId"))
	if err != nil {
		response.APIResponse(c, err.Error(), 404, data)
		return
	}

	response.APIResponse(c, data, 200, "CSV file path")
}

func (h *BaseHandler) GetUserResult(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()

	email := c.GetHeader("Email")
	data, err := h.BaseSvc.FetchUserResult(ctx, c.Param("moduleId"), email)
	if err != nil {
		response.APIResponse(c, err.Error(), 404, data.(string))
		return
	}

	response.APIResponse(c, data, 200, "User test result")
}
