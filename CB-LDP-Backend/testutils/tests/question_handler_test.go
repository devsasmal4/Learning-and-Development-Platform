package tests

import (
	"cb-ldp-backend/testutils/mocks"
	"net/http"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (t *BaseHandlerTestSuite) TestQuestionHandler_GetAnswer_ShouldHandleSuccess() {
	question := mocks.MockQuestion()
	option := mocks.MockOption()
	
	t.engine.GET("/question/:questionId/answer/:optionId", t.handler.GetAnswer)
	t.mockBaseService.EXPECT().FetchCorrectAnswer(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(option.OptionString, nil)
	url := "/question/" + question.Id.Hex() + "/answer/" + option.Id.Hex()
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	expectedResponse := `{"status":200,"message":"Correct answer","response_data":"option1"}`
	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestQuestionHandler_GetAnswer_ShouldHandleBadRequest() {
	t.engine.GET("/question/:questionId/answer/:optionId", t.handler.GetAnswer)
	t.mockBaseService.EXPECT().FetchCorrectAnswer(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return("", nil)
	url := "/question/" + "/answer/"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 404, t.recorder.Code)
}

func (t *BaseHandlerTestSuite) TestQuestionHandler_DeleteQuestion_ShouldHandleSuccess() {
	question := mocks.MockQuestion()
	t.engine.GET("/question/:questionId", t.handler.DeleteQuestion)
	t.mockBaseService.EXPECT().DeleteQuestion(gomock.Any(), gomock.Any()).AnyTimes().Return(" Question Deleted", 200, nil)
	url := "/question/" + question.Id.Hex()
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	expectedResponse := `{"status":200,"message":" Question Deleted"}`
	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestQuestionHandler_DeleteQuestion_ShouldHandleBadRequest() {
	t.engine.GET("/question/:questionId", t.handler.DeleteQuestion)
	t.mockBaseService.EXPECT().DeleteQuestion(gomock.Any(), gomock.Any()).AnyTimes().Return(" Question Deleted", 200, nil)
	url := "/question/"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 404, t.recorder.Code)
}
