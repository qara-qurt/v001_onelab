package http

import (
	"bytes"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"net/http/httptest"
	"testing"
	"v001_onelab/internal/model"
	"v001_onelab/internal/service"
	mock_service "v001_onelab/internal/service/mocks"
)

func TestHandler_SignUp(t *testing.T) {

	type mockBehavior func(s *mock_service.MockIUser, book model.UserInput)

	testTable := []struct {
		name               string
		inputBody          string
		inputUser          model.UserInput
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"fullname":"test-name","login":"test-login","password":"test-password"}`,
			inputUser: model.UserInput{
				FullName: "test-name",
				Login:    "test-login",
				Password: "test-password",
			},
			mockBehavior: func(s *mock_service.MockIUser, user model.UserInput) {
				s.EXPECT().Create(user)
			},
			expectedStatusCode: 201,
		},
		{
			name:      "NO REQUIRED FIELDS",
			inputBody: `{"fullname":"test-name","login":"test-login"}`,
			mockBehavior: func(s *mock_service.MockIUser, user model.UserInput) {
			},
			expectedStatusCode: 400,
		},
		{
			name:      "VALIDATE ERROR",
			inputBody: `{"fullname":"test-name","login":"","password":""}`,
			mockBehavior: func(s *mock_service.MockIUser, user model.UserInput) {
			},
			expectedStatusCode: 400,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_service.NewMockIUser(c)
			testCase.mockBehavior(user, testCase.inputUser)

			services := &service.Service{User: user}
			handler := New(services)

			h := echo.New()
			h.POST("/api/auth/sign-up", handler.SignUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/auth/sign-up", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Content-Type", "application/json")

			h.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
