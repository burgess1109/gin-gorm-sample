package message

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gin-gorm-sample/mocks/application/message/port"
)

func TestPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var testCases = []struct {
		testName               string
		requestBody            io.Reader
		applicationReturnError error
		expectedStatusCode     int
		expectedResponse       string
	}{
		{
			"should be failed if binding error",
			bytes.NewBufferString(""),
			nil,
			http.StatusBadRequest,
			`{"error":"EOF"}`,
		},
		{
			"should be failed if no context",
			bytes.NewBufferString(`{"user_id": 123}`),
			nil,
			http.StatusBadRequest,
			`{"error":"Key: 'postRequestBody.Context' Error:Field validation for 'Context' failed on the 'required' tag"}`,
		},
		{
			"should be failed if no user_id",
			bytes.NewBufferString(`{"context":"TEST"}`),
			nil,
			http.StatusBadRequest,
			`{"error":"Key: 'postRequestBody.UserID' Error:Field validation for 'UserID' failed on the 'required' tag"}`,
		},
		{
			"should be failed if application return error",
			bytes.NewBufferString(`{"context":"TEST", "user_id": 123}`),
			errors.New("application error"),
			http.StatusInternalServerError,
			`{"error":"application error"}`,
		},
		{
			"should be pass",
			bytes.NewBufferString(`{"context":"TEST", "user_id": 123}`),
			nil,
			http.StatusOK,
			`{"message":"create user success"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			mockPort := port.NewMockWeb(ctrl)
			router := NewRouter(mockPort)

			mockPort.EXPECT().Create(gomock.Any()).Return(testCase.applicationReturnError).AnyTimes()

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("POST", "messages", testCase.requestBody)

			router.post(c)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponse+"\n", w.Body.String())
		})
	}
}
