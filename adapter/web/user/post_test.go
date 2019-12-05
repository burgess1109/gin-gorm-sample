package user

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

	"gin-gorm-sample/mocks/application/user/port"
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
			"should be failed if no name",
			bytes.NewBufferString(`{"sex": "boy", "email": "abc@gmail.com"}`),
			nil,
			http.StatusBadRequest,
			`{"error":"Key: 'postRequestBody.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`,
		},
		{
			"should be failed if length of name greater than 50",
			bytes.NewBufferString(`{"name":"TEST:this_is_a_string_which_is_greater_than_50_length", "sex": "boy", "email": "abc@gmail.com"}`),
			nil,
			http.StatusBadRequest,
			`{"error":"Key: 'postRequestBody.Name' Error:Field validation for 'Name' failed on the 'lt' tag"}`,
		},
		{
			"should be failed if no sex",
			bytes.NewBufferString(`{"name":"TEST", "email": "abc@gmail.com"}`),
			nil,
			http.StatusBadRequest,
			`{"error":"Key: 'postRequestBody.Sex' Error:Field validation for 'Sex' failed on the 'required' tag"}`,
		},
		{
			"should be failed if no email",
			bytes.NewBufferString(`{"name":"TEST", "sex": "boy"}`),
			nil,
			http.StatusBadRequest,
			`{"error":"Key: 'postRequestBody.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
		},
		{
			"should be failed if wrong email format",
			bytes.NewBufferString(`{"name":"TEST", "sex": "boy", "email":"aaa.bbb.ccc"}`),
			nil,
			http.StatusBadRequest,
			`{"error":"Key: 'postRequestBody.Email' Error:Field validation for 'Email' failed on the 'email' tag"}`,
		},
		{
			"should be failed if application return error",
			bytes.NewBufferString(`{"name":"TEST", "sex": "boy", "email": "abc@gmail.com"}`),
			errors.New("application error"),
			http.StatusInternalServerError,
			`{"error":"application error"}`,
		},
		{
			"should be pass",
			bytes.NewBufferString(`{"name":"TEST", "sex": "boy", "email": "abc@gmail.com"}`),
			nil,
			http.StatusOK,
			`{"message":"create user successful"}`,
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
			c.Request, _ = http.NewRequest("POST", "users", testCase.requestBody)

			router.post(c)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponse+"\n", w.Body.String())
		})
	}
}
