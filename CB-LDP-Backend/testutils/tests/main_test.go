package tests

import (
	"cb-ldp-backend/handlers"
	"cb-ldp-backend/testutils/mocks"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type BaseHandlerTestSuite struct {
	suite.Suite
	mockCtrl        *gomock.Controller
	context         *gin.Context
	engine          *gin.Engine
	recorder        *httptest.ResponseRecorder
	mockBaseService *mocks.MockIBaseService
	handler         handlers.IBaseHandler
}

func TestBaseHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(BaseHandlerTestSuite))
}

func (t *BaseHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	t.mockCtrl = gomock.NewController(t.T())
	t.recorder = httptest.NewRecorder()
	t.context, t.engine = gin.CreateTestContext(t.recorder)
	t.mockBaseService = mocks.NewMockIBaseService(t.mockCtrl)
	t.handler = handlers.NewBaseHandler(t.mockBaseService)
}

func (t *BaseHandlerTestSuite) TearDownTest() {
	t.mockCtrl.Finish()
	t.recorder.Flush()
}
