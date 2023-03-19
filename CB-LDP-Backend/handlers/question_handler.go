package handlers

import (
	"cb-ldp-backend/commons/utility"
	"cb-ldp-backend/config"
	"cb-ldp-backend/constants"
	"cb-ldp-backend/models/response"
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

var envVar = config.LoadConfig()

// func (h *BaseHandler) UpdateQuestion(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
// 	defer cancel()

// 	var question entity.Question
// 	if err := json.NewDecoder(c.Request.Body).Decode(&question); err != nil {
// 		response.APIResponse(c, err.Error(), 400, "Bad Request")
// 		return
// 	}

// 	var reqBody entity.RequestBody = question
// 	if err := reqBody.ValidateStruct(); err != "" {
// 		response.APIResponse(c, err, 400, "Bad Request")
// 		return
// 	}
// 	err := h.BaseSvc.UpdateQuestion(ctx, c.Param("questionId"), question)
// 	if err != nil {
// 		response.APIResponse(c, "Question not found", 404, "Bad Request")
// 		return
// 	}

// 	response.APIResponse(c, nil, 200, "Question updated")
// }

// @Summary      Delete question
// @Description  Delete question
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param questionId path string true "Question Id"
// @Success      200  {object}  response.Response
// @failure      404              {string}    "Question not found"
// @failure      500              {string}    "Error deleting options"
// @Router       /question/{questionId} [delete]
func (h *BaseHandler) DeleteQuestion(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	data, code, err := h.BaseSvc.DeleteQuestion(ctx, c.Param("questionId"))

	if err != nil {
		response.APIResponse(c, err.Error(), code, data)
		return
	}
	response.APIResponse(c, nil, code, data)

}

// @Summary      Get correct answer
// @Description  Get correct answer
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param questionId path string true "Question Id"
// @Param optionId path string true "Option Id"
// @Success      200  {object}  response.AnswerResponse
// @failure      404              {string}    "Question not found"
// @failure      500              {string}    "Server Error"
// @Router       /question/{questionId}/answer/{optionId} [get]
func (h *BaseHandler) GetAnswer(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	data, err := h.BaseSvc.FetchCorrectAnswer(ctx, c.Param("questionId"), c.Param("optionId"))
	if err != nil {
		response.APIResponse(c, err.Error(), 404, data.(string))
		return
	}
	response.APIResponse(c, data, 200, "Correct answer")
}

// @Summary      Upload CSV
// @Description  Upload CSV
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param file  formData file true  "Questionaires in CSV format"
// @Success      200  {object}  response.QuestionJsonResponse
// @failure      400              {string}    "Bad request"
// @failure      500              {string}    "Error creating CSV"
// @Router       /question/upload [post]
func (h *BaseHandler) UploadCsv(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.APIResponse(c, err.Error(), 400, "Bad request")
		return
	}
	dirPath := envVar["dir_path"].(string) + "uploads/"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		response.APIResponse(c, err.Error(), 500, "Error creating CSV")
		return
	}

	fileName := dirPath + uuid.NewV4().String() + "_" + file.Filename
	c.SaveUploadedFile(file, fileName)
	questions, err := utility.ConvertCsvToJson(fileName)
	if err != nil {
		response.APIResponse(c, err.Error(), 500, "Error creating CSV")
		return
	}

	err = os.Remove(fileName)
	if err != nil {
		response.APIResponse(c, err.Error(), 500, "Error deleting CSV")
		return
	}

	response.APIResponse(c, questions, 200, "Questions uploaded successfully")
}
