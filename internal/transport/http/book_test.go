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

func TestHandler_CreateBook(t *testing.T) {
	type mockBehavior func(s *mock_service.MockIBook, book model.BookInput)

	testTable := []struct {
		name               string
		inputBody          string
		inputBook          model.BookInput
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: `{"name":"test-name","description":"test-description","author":"test-author"}`,
			inputBook: model.BookInput{
				Name:        "test-name",
				Description: "test-description",
				Author:      "test-author",
			},
			mockBehavior: func(s *mock_service.MockIBook, book model.BookInput) {
				s.EXPECT().Create(book)
			},
			expectedStatusCode: 201,
		},
		{
			name:      "NO REQUIRED FIELDS",
			inputBody: `{"description":"test-description","author":"test-author"}`,
			mockBehavior: func(s *mock_service.MockIBook, book model.BookInput) {
			},
			expectedStatusCode: 400,
		},
		{
			name:      "VALIDATE ERROR",
			inputBody: `{"name":"test-name","description":"","author":""}`,
			mockBehavior: func(s *mock_service.MockIBook, book model.BookInput) {
			},
			expectedStatusCode: 400,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			book := mock_service.NewMockIBook(c)
			testCase.mockBehavior(book, testCase.inputBook)

			services := &service.Service{Book: book}
			handler := New(services)

			h := echo.New()
			h.POST("/api/books/", handler.CreateBook)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/books/", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("Content-Type", "application/json")

			h.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
