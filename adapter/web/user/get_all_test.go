package user

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
	"gin-gorm-sample/mocks/application/user/port"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	responseBodyJSON, _ := json.Marshal(GetAllResponseBody{Data: []User{}})

	var testCases = []struct {
		testName               string
		applicationReturnError error
		expectedStatusCode     int
		expectedResponse       string
	}{
		{
			"should be failed if application return error",
			errors.New("application error"),
			http.StatusInternalServerError,
			`{"error":"application error"}`,
		},
		{
			"should be pass",
			nil,
			http.StatusOK,
			string(responseBodyJSON),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			mockPort := port.NewMockWeb(ctrl)
			router := NewRouter(mockPort)

			mockPort.EXPECT().GetAll().Return([]domain.User{}, testCase.applicationReturnError).AnyTimes()

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("GET", "users", nil)

			router.getAll(c)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponse+"\n", w.Body.String())
		})
	}
}
