package tests

import (
	"bytes"
	"cb-ldp-backend/testutils/mocks"
	"encoding/json"
	"net/http"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


func (t *BaseHandlerTestSuite) TestQuizHandler_ViewQuiz_ShouldHandleSuccess() {
	quiz := mocks.MockModule().ModuleQuiz
	t.engine.GET("/quiz/:moduleId", t.handler.ViewQuiz)
	t.mockBaseService.EXPECT().ViewQuiz(gomock.Any(), gomock.Any()).AnyTimes().Return(quiz.QuizName, nil)
	url := "/quiz/" + mocks.MockModule().Id.Hex()
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	expectedResponse := `{"status":200,"message":"Questionaires","response_data":"name1"}`
	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestQuizHandler_ViewQuiz_ShouldHandleBadRequest() {
	quiz := mocks.MockModule().ModuleQuiz
	t.engine.GET("/quiz/:moduleId", t.handler.ViewQuiz)
	t.mockBaseService.EXPECT().ViewQuiz(gomock.Any(), gomock.Any()).AnyTimes().Return(quiz.QuizName, nil)
	url := "/quiz/"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 404, t.recorder.Code)
}

func (t *BaseHandlerTestSuite) TestQuizHandler_ExecuteQuiz_ShouldHandleSuccess() {
	quizResponse := mocks.MockQuizResponse()
	t.engine.GET("/quiz/execute", t.handler.ExecuteQuiz)
	t.mockBaseService.EXPECT().CreateQuizResponse(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return("Quiz executed sucessfully", 200, nil)
	url := "/quiz/execute"
	
	jsonValue, _ := json.Marshal(quizResponse)
	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Email", "test@coffeebeans.io")
	expectedResponse := `{"status":200,"message":"Quiz executed successfully","response_data":"Quiz executed sucessfully"}`
	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestQuizHandler_ExecuteQuiz_ShouldHandleBadRequest() {
	t.engine.GET("/quiz/execute", t.handler.ExecuteQuiz)
	t.mockBaseService.EXPECT().CreateQuizResponse(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return("Quiz executed sucessfully", 200, nil)
	url := "/quiz/execute"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	expectedResponse := `{"status":400,"message":"Quiz execution failed!","response_data":"invalid request"}`
	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 400, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}