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

func TestPatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var testCases = []struct {
		testName               string
		requestID              string
		requestBody            io.Reader
		applicationReturnError error
		expectedStatusCode     int
		expectedResponse       string
	}{
		{
			"should be failed if binding error",
			"123",
			bytes.NewBufferString(""),
			nil,
			http.StatusBadRequest,
			`{"error":"EOF"}`,
		},
		{
			"should be failed if id is not a number",
			"abc",
			bytes.NewBufferString("{}"),
			nil,
			http.StatusBadRequest,
			`{"error":"strconv.Atoi: parsing \"abc\": invalid syntax"}`,
		},
		{
			"should be failed if application return error",
			"123",
			bytes.NewBufferString("{}"),
			errors.New("application error"),
			http.StatusInternalServerError,
			`{"error":"application error"}`,
		},
		{
			"should be pass",
			"123",
			bytes.NewBufferString("{}"),
			nil,
			http.StatusOK,
			`{"message":"update user successful"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			mockPort := port.NewMockWeb(ctrl)
			router := NewRouter(mockPort)

			mockPort.EXPECT().Update(gomock.Any()).Return(testCase.applicationReturnError).AnyTimes()

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{gin.Param{Key: "id", Value: testCase.requestID}}
			c.Request, _ = http.NewRequest("PATCH", "messages", testCase.requestBody)

			router.patch(c)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponse+"\n", w.Body.String())
		})
	}
}
