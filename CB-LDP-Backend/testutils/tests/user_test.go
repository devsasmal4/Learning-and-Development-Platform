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

func (t *BaseHandlerTestSuite) TestCreateUserToken_ShouldHandleSuccess() {
	t.engine.POST("/user/createToken", t.handler.GenerateToken)
	user := mocks.MockUser()

	t.mockBaseService.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/user/createToken", bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
}

func (t *BaseHandlerTestSuite) TestUpdateUserTokenShouldHandleSuccess() {
	t.engine.POST("/user/createToken", t.handler.GenerateToken)
	user := mocks.MockUser()
	responsBody := `{"status":200,"message":"User Token Created"}`

	t.mockBaseService.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/user/createToken", bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 200, t.recorder.Code)
	assert.Equal(t.T(), responsBody, t.recorder.Body.String())

}

func (t *BaseHandlerTestSuite) TestTokenCreation_ShouldHandleFailure() {
	t.engine.POST("/user/createToken", t.handler.GenerateToken)
	user := entity.User{
		UserName: "Dev",
		UserMail: "deva@gmail.com",
	}

	t.mockBaseService.EXPECT().GenerateToken(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/user/createToken", bytes.NewBuffer(jsonValue))

	t.engine.ServeHTTP(t.recorder, req)
	assert.Equal(t.T(), 400, t.recorder.Code)
}
