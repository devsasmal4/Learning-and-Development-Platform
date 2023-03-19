package handlers

import (
	"cb-ldp-backend/constants"
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/models/response"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// func (h *BaseHandler) CreateQuiz(c *gin.Context) {
// 	var quiz entity.Quiz
// 	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
// 	defer cancel()

// 	if err := json.NewDecoder(c.Request.Body).Decode(&quiz); err != nil {
// 		response.APIResponse(c, err.Error(), 400, "Bad Request")
// 		return
// 	}
// 	var reqBody entity.RequestBody = quiz
// 	if err := reqBody.ValidateStruct(); err != "" {
// 		response.APIResponse(c, err, 400, "Bad Request")
// 		return
// 	}

// 	err := h.BaseSvc.CreateQuiz(ctx)
// 	if err != nil {
// 		response.APIResponse(c, err.Error(), 500, "Error creating quiz")
// 		return
// 	}
// 	response.APIResponse(c, nil, 200, "Successfully created quiz")
// }

// @Summary      Update quiz
// @Description  Update quiz
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param quizId path string true "Quiz Id"
// @Param entity.Quiz  body entity.Quiz true  "Updated quiz request body"
// @Success      200  {object}  response.Response
// @failure      400              {string}    "Bad Request"
// @failure      404              {string}    "Quiz not found"
// @Router       /quiz/{quizId} [put]
// func (h *BaseHandler) UpdateQuiz(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
// 	defer cancel()

// 	var quiz entity.Quiz
// 	if err := json.NewDecoder(c.Request.Body).Decode(&quiz); err != nil {
// 		response.APIResponse(c, err.Error(), 400, "Bad Request")
// 		return
// 	}

// 	var reqBody entity.RequestBody = quiz
// 	if err := reqBody.ValidateStruct(); err != "" {
// 		response.APIResponse(c, err, 400, "Bad Request")
// 		return
// 	}
// 	err := h.BaseSvc.UpdateQuiz(ctx, c.Param("quizId"))
// 	if err != nil {
// 		response.APIResponse(c, "Quiz not found", 404, "Bad Request")
// 		return
// 	}
// 	response.APIResponse(c, nil, 200, "Quiz updated")

// }

// func (h *BaseHandler) DeleteQuiz(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
// 	defer cancel()

// 	err := h.BaseSvc.DeleteQuiz(ctx, c.Param("quizId"))
// 	if err != nil {
// 		response.APIResponse(c, err.Error(), 500, "Error deleting questions and options")
// 		return
// 	}
// 	response.APIResponse(c, nil, 200, "Successfully deleted quiz")

// }

// @Summary      View quiz
// @Description  View quiz
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param moduleId path string true "Module Id"
// @Success      200  {object}  response.QuizQuestions
// @failure      404              {string}    "Module not found"
// @failure      500              {string}    "Server Error"
// @Router       /quiz/{moduleId} [get]
func (h *BaseHandler) ViewQuiz(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	quizQuestions, err := h.BaseSvc.ViewQuiz(ctx, c.Param("moduleId"))

	if err != nil {
		response.APIResponse(c, err.Error(), 404, "Module not found")
		return
	}
	response.APIResponse(c, quizQuestions, 200, "Questionaires")

}

// @Summary     Quiz execution
// @Description  Quiz execution
// @Produce      json
// @Param Email  header string true "Email Id"
// @Param token  header string true "Token"
// @Param entity.QuizResponse  body entity.QuizResponse true  "Quiz response body"
// @Success      200  {object}  entity.QuizResponse
// @failure      404              {string}    "Module not found"
// @failure      500              {string}    "Server Error"
// @Router       /quiz/execute [post]
func (h *BaseHandler) ExecuteQuiz(c *gin.Context) {
	var quizResponse entity.QuizResponse
	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultTimeOut*time.Second)
	defer cancel()
	email := c.GetHeader("Email")
	if err := c.ShouldBindJSON(&quizResponse); err != nil {
		response.APIResponse(c, err.Error(), 400, "Quiz execution failed!")
		return
	}

	data, code, err := h.BaseSvc.CreateQuizResponse(ctx, email, quizResponse)

	if err != nil {
		response.APIResponse(c, err.Error(), code, data.(string))
		return
	}
	response.APIResponse(c, data, code, "Quiz executed successfully")

}
