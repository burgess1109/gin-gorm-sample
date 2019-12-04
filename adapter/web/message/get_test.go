package message

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"gin-gorm-sample/application/domain"
	"gin-gorm-sample/mocks/application/message/port"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	responseBodyJSON, _ := json.Marshal(GetResponseBody{})

	var testCases = []struct {
		testName               string
		requestID              string
		applicationReturnError error
		expectedStatusCode     int
		expectedResponse       string
	}{
		{
			"should be failed if id is not a number",
			"abc",
			nil,
			http.StatusBadRequest,
			`{"error":"strconv.Atoi: parsing \"abc\": invalid syntax"}`,
		},
		{
			"should be failed if application return error",
			"123",
			errors.New("application error"),
			http.StatusInternalServerError,
			`{"error":"application error"}`,
		},
		{
			"should be pass",
			"123",
			nil,
			http.StatusOK,
			string(responseBodyJSON),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			mockPort := port.NewMockWeb(ctrl)
			router := NewRouter(mockPort)

			mockPort.EXPECT().Get(gomock.Any()).Return(domain.Message{}, testCase.applicationReturnError).AnyTimes()

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{gin.Param{Key: "id", Value: testCase.requestID}}
			c.Request, _ = http.NewRequest("GET", "messages", nil)

			router.get(c)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponse+"\n", w.Body.String())
		})
	}
}
