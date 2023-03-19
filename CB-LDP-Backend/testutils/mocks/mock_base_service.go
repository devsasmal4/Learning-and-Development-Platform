// Code generated by MockGen. DO NOT EDIT.
// Source: base_service.go

// Package mock_service is a generated GoMock package.
package mocks

import (
	entity "cb-ldp-backend/models/entity"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIBaseService is a mock of IBaseService interface.
type MockIBaseService struct {
	ctrl     *gomock.Controller
	recorder *MockIBaseServiceMockRecorder
}

// MockIBaseServiceMockRecorder is the mock recorder for MockIBaseService.
type MockIBaseServiceMockRecorder struct {
	mock *MockIBaseService
}

// NewMockIBaseService creates a new mock instance.
func NewMockIBaseService(ctrl *gomock.Controller) *MockIBaseService {
	mock := &MockIBaseService{ctrl: ctrl}
	mock.recorder = &MockIBaseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBaseService) EXPECT() *MockIBaseServiceMockRecorder {
	return m.recorder
}

// CreateModule mocks base method.
func (m *MockIBaseService) CreateModule(arg0 context.Context, arg1 *entity.Module) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateModule", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateModule indicates an expected call of CreateModule.
func (mr *MockIBaseServiceMockRecorder) CreateModule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModule", reflect.TypeOf((*MockIBaseService)(nil).CreateModule), arg0, arg1)
}

// CreateQuizResponse mocks base method.
func (m *MockIBaseService) CreateQuizResponse(ctx context.Context, email string, quizResponse entity.QuizResponse) (interface{}, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuizResponse", ctx, email, quizResponse)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateQuizResponse indicates an expected call of CreateQuizResponse.
func (mr *MockIBaseServiceMockRecorder) CreateQuizResponse(ctx, email, quizResponse interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuizResponse", reflect.TypeOf((*MockIBaseService)(nil).CreateQuizResponse), ctx, email, quizResponse)
}

// DeleteModule mocks base method.
func (m *MockIBaseService) DeleteModule(ctx context.Context, moduleId string) (int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteModule", ctx, moduleId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteModule indicates an expected call of DeleteModule.
func (mr *MockIBaseServiceMockRecorder) DeleteModule(ctx, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteModule", reflect.TypeOf((*MockIBaseService)(nil).DeleteModule), ctx, moduleId)
}

// DeleteQuestion mocks base method.
func (m *MockIBaseService) DeleteQuestion(ctx context.Context, questionId string) (string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQuestion", ctx, questionId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteQuestion indicates an expected call of DeleteQuestion.
func (mr *MockIBaseServiceMockRecorder) DeleteQuestion(ctx, questionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuestion", reflect.TypeOf((*MockIBaseService)(nil).DeleteQuestion), ctx, questionId)
}

// FetchAllModules mocks base method.
func (m *MockIBaseService) FetchAllModules(ctx context.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllModules", ctx)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllModules indicates an expected call of FetchAllModules.
func (mr *MockIBaseServiceMockRecorder) FetchAllModules(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllModules", reflect.TypeOf((*MockIBaseService)(nil).FetchAllModules), ctx)
}

// FetchCorrectAnswer mocks base method.
func (m *MockIBaseService) FetchCorrectAnswer(ctx context.Context, questionId, optionId string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCorrectAnswer", ctx, questionId, optionId)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCorrectAnswer indicates an expected call of FetchCorrectAnswer.
func (mr *MockIBaseServiceMockRecorder) FetchCorrectAnswer(ctx, questionId, optionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCorrectAnswer", reflect.TypeOf((*MockIBaseService)(nil).FetchCorrectAnswer), ctx, questionId, optionId)
}

// FetchModule mocks base method.
func (m *MockIBaseService) FetchModule(ctx context.Context, moduleId string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchModule", ctx, moduleId)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchModule indicates an expected call of FetchModule.
func (mr *MockIBaseServiceMockRecorder) FetchModule(ctx, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchModule", reflect.TypeOf((*MockIBaseService)(nil).FetchModule), ctx, moduleId)
}

// FetchModuleInstruction mocks base method.
func (m *MockIBaseService) FetchModuleInstruction(ctx context.Context, moduleId string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchModuleInstruction", ctx, moduleId)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchModuleInstruction indicates an expected call of FetchModuleInstruction.
func (mr *MockIBaseServiceMockRecorder) FetchModuleInstruction(ctx, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchModuleInstruction", reflect.TypeOf((*MockIBaseService)(nil).FetchModuleInstruction), ctx, moduleId)
}

// FetchTestDetails mocks base method.
func (m *MockIBaseService) FetchTestDetails(ctx context.Context, moduleId string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTestDetails", ctx, moduleId)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTestDetails indicates an expected call of FetchTestDetails.
func (mr *MockIBaseServiceMockRecorder) FetchTestDetails(ctx, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTestDetails", reflect.TypeOf((*MockIBaseService)(nil).FetchTestDetails), ctx, moduleId)
}

// GenerateCsv mocks base method.
func (m *MockIBaseService) GenerateCsv(ctx context.Context, moduleId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCsv", ctx, moduleId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCsv indicates an expected call of GenerateCsv.
func (mr *MockIBaseServiceMockRecorder) GenerateCsv(ctx, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCsv", reflect.TypeOf((*MockIBaseService)(nil).GenerateCsv), ctx, moduleId)
}

// GenerateToken mocks base method.
func (m *MockIBaseService) GenerateToken(ctx context.Context, user entity.User) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, user)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockIBaseServiceMockRecorder) GenerateToken(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockIBaseService)(nil).GenerateToken), ctx, user)
}

// UpdateModule mocks base method.
func (m *MockIBaseService) UpdateModule(ctx context.Context, module entity.Module, moduleId string) (string, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateModule", ctx, module, moduleId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateModule indicates an expected call of UpdateModule.
func (mr *MockIBaseServiceMockRecorder) UpdateModule(ctx, module, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateModule", reflect.TypeOf((*MockIBaseService)(nil).UpdateModule), ctx, module, moduleId)
}

// ViewQuiz mocks base method.
func (m *MockIBaseService) ViewQuiz(ctx context.Context, moduleId string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewQuiz", ctx, moduleId)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewQuiz indicates an expected call of ViewQuiz.
func (mr *MockIBaseServiceMockRecorder) ViewQuiz(ctx, moduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewQuiz", reflect.TypeOf((*MockIBaseService)(nil).ViewQuiz), ctx, moduleId)
}
