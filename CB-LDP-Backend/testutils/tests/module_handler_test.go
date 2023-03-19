package tests

import (
	"bytes"
	"cb-ldp-backend/models/entity"
	"cb-ldp-backend/testutils/mocks"
	"encoding/json"

	"net/http"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (t *BaseHandlerTestSuite) TestModuleHandler_CreateModule_ShouldHandleSuccess() {
	t.engine.POST("/module/", t.handler.CreateModule)
	module := mocks.MockModule()
	expectedResponse := `{"status":200,"message":"Module Created"}`

	t.mockBaseService.EXPECT().CreateModule(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	jsonValue, _ := json.Marshal(module)
	req, _ := http.NewRequest(http.MethodPost, "/module/", bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestModuleHandler_CreateModule_ShouldHandleValidationFail() {
	t.engine.POST("/module/", t.handler.CreateModule)
	instrctions := [...]string{"Instruction 1", "Instruction 2", "Instruction 3"}
	module := entity.Module{
		ModuleName:          "1",
		ModuleStudyMaterial: "Link",
		ModuleCreatedBy:     "User 1",
		ModuleDepartment:    "a",
		ModuleInstructions:  instrctions[:],
		ModuleDescription:   "a",
		ModuleStatus:        true,
		ModuleQuiz: entity.Quiz{
			QuizName:            "name1",
			QuizPassingMarks:    50,
			QuizDurationMinutes: 90,
			QuizTotalMarks:      100,
			QuizQuestionaires:   []entity.Question{},
		},
	}
	t.mockBaseService.EXPECT().CreateModule(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	jsonValue, _ := json.Marshal(module)
	req, _ := http.NewRequest(http.MethodPost, "/module/", bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 400, t.recorder.Code)
}

func (t *BaseHandlerTestSuite) TestModuleHandler_ViewModule_ShouldHandleSucess() {
	module := mocks.MockModule()
	t.engine.GET("/module/:modulId", t.handler.ViewModule)
	expectedResponse := `{"status":200,"message":"Module displayed","response_data":"Module 1"}`
	t.mockBaseService.EXPECT().FetchModule(gomock.Any(), gomock.Any()).AnyTimes().Return(module.ModuleName, nil)
	url := "/module/" + module.Id.Hex()
	req, _ := http.NewRequest(http.MethodGet, url+module.Id.Hex(), nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestModuleHandler_ViewModule_ShouldHandleBadRequest() {
	module := mocks.MockModule()
	t.engine.GET("/module/:modulId", t.handler.ViewModule)
	t.mockBaseService.EXPECT().FetchModule(gomock.Any(), gomock.Any()).AnyTimes().Return(module, nil)
	url := "/module/"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 404, t.recorder.Code)
}

func (t *BaseHandlerTestSuite) TestModuleHandler_UpdateModule_ShouldHandleSucess() {
	module := mocks.MockModule()
	module.ModuleName = "No name module!"
	expectedResponse := `{"status":200,"message":"Module Updated"}`
	t.engine.GET("/module/:modulId", t.handler.UpdateModule)
	t.mockBaseService.EXPECT().UpdateModule(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return("Module Updated", 200, nil)
	url := "/module/" + module.Id.Hex()
	jsonValue, _ := json.Marshal(module)
	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestModuleHandler_UpdateModule_ShouldHandleBadRequest() {
	t.engine.GET("/module/:modulId", t.handler.UpdateModule)
	t.mockBaseService.EXPECT().UpdateModule(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return("Module Updated", 200, nil)
	url := "/module/"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 404, t.recorder.Code)
}

func (t *BaseHandlerTestSuite) TestModuleHandler_DeleteModule_ShouldHandleSucess() {
	module := mocks.MockModule()
	expectedResponse := `{"status":200,"message":"Module Deleted"}`
	t.engine.GET("/module/:modulId", t.handler.DeleteModule)
	t.mockBaseService.EXPECT().DeleteModule(gomock.Any(), gomock.Any()).AnyTimes().Return(200, "Module Deleted", nil)
	url := "/module/" + module.Id.Hex()
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestModuleHandler_DeleteModule_ShouldHandleBadRequest() {
	t.engine.GET("/module/:modulId", t.handler.DeleteModule)
	t.mockBaseService.EXPECT().DeleteModule(gomock.Any(), gomock.Any()).AnyTimes().Return(200, "Module Deleted", nil)
	url := "/module/"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 404, t.recorder.Code)
}

func (t *BaseHandlerTestSuite) TestModuleHandler_TestDetails_ShouldHandleSucess() {
	module := mocks.MockModule()
	testDetails := mocks.MockTestDetailsResponse()
	expectedResponse := `{"status":200,"message":"Details Displayed","response_data":[{"id":"635fb6560fc5f5c7d38976ec","user_name":"user","module_name":"Module 1","start_date":"0001-01-01T00:00:00Z","end_date":"0001-01-01T00:00:00.000000003Z","test_status":"Completed","quiz_score":70,"quiz_result":true,"duration":"3ns"}]}`

	t.engine.GET("/module/:modulId/testDetails", t.handler.GetTestDetails)
	t.mockBaseService.EXPECT().FetchTestDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(testDetails, nil)
	url := "/module/" + module.Id.Hex() + "/testDetails"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), expectedResponse, t.recorder.Body.String())
}

func (t *BaseHandlerTestSuite) TestModuleHandler_TestDetails_ShouldHandleBadRequest() {
	testDetails := mocks.MockTestDetailsResponse()
	t.engine.GET("/module/:modulId/testDetails", t.handler.GetTestDetails)
	t.mockBaseService.EXPECT().FetchTestDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(testDetails, nil)
	url := "/module/testDetails"
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 404, t.recorder.Code)
}
