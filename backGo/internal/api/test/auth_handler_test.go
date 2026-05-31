package api_test

import (
	"bytes"
	"boock/backGo/internal/api"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Note 1: MockAuthService 는 AuthServiceInterface 를 구현하는 Mock 객체입니다.
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(name, password string) error {
	args := m.Called(name, password)
	return args.Error(0)
}

func (m *MockAuthService) Login(congCode, email, password string) (string, error) {
	args := m.Called(congCode, email, password)
	return args.String(0), args.Error(1)
}

func TestAuthHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Register", func(t *testing.T) {
		tests := []struct {
			name           string
			input          interface{}
			mockFn         func(m *MockAuthService)
			expectedStatus int
			expectedBody   string
		}{
			{
				name: "성공적인 회원가입",
				input: gin.H{
					"name":     "testuser",
					"password": "password123",
				},
				mockFn: func(m *MockAuthService) {
					m.On("Register", "testuser", "password123").Return(nil)
				},
				expectedStatus: http.StatusCreated,
				expectedBody:   "가입 성공",
			},
			{
				name: "잘못된 입력값 (필드 누락)",
				input: gin.H{
					"name": "testuser",
				},
				mockFn:         func(m *MockAuthService) {},
				expectedStatus: http.StatusBadRequest,
				expectedBody:   "입력값을 확인해주세요.",
			},
			{
				name: "서비스 레이어 에러",
				input: gin.H{
					"name":     "testuser",
					"password": "password123",
				},
				mockFn: func(m *MockAuthService) {
					m.On("Register", "testuser", "password123").Return(errors.New("db error"))
				},
				expectedStatus: http.StatusInternalServerError,
				expectedBody:   "회원가입에 실패했습니다.",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockService := new(MockAuthService)
				tt.mockFn(mockService)
				handler := api.NewAuthHandler(mockService)

				w := httptest.NewRecorder()
				_, r := gin.CreateTestContext(w)
				r.POST("/register", handler.Register)

				jsonInput, _ := json.Marshal(tt.input)
				req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonInput))
				req.Header.Set("Content-Type", "application/json")

				r.ServeHTTP(w, req)

				assert.Equal(t, tt.expectedStatus, w.Code)
				assert.Contains(t, w.Body.String(), tt.expectedBody)
				mockService.AssertExpectations(t)
			})
		}
	})

	t.Run("Login", func(t *testing.T) {
		tests := []struct {
			name           string
			input          interface{}
			mockFn         func(m *MockAuthService)
			expectedStatus int
			expectedBody   string
		}{
			{
				name: "성공적인 로그인",
				input: gin.H{
					"cong_code":    "123",
					"jwhub_email":  "test@example.com",
					"password":     "password123",
				},
				mockFn: func(m *MockAuthService) {
					m.On("Login", "123", "test@example.com", "password123").Return("fake-jwt-token", nil)
				},
				expectedStatus: http.StatusOK,
				expectedBody:   "fake-jwt-token",
			},
			{
				name: "잘못된 입력값",
				input: gin.H{
					"cong_code": "123",
				},
				mockFn:         func(m *MockAuthService) {},
				expectedStatus: http.StatusBadRequest,
				expectedBody:   "입력값을 확인해주세요.",
			},
			{
				name: "인증 실패 (잘못된 비밀번호 등)",
				input: gin.H{
					"cong_code":    "123",
					"jwhub_email":  "test@example.com",
					"password":     "wrong",
				},
				mockFn: func(m *MockAuthService) {
					m.On("Login", "123", "test@example.com", "wrong").Return("", errors.New("unauthorized"))
				},
				expectedStatus: http.StatusUnauthorized,
				expectedBody:   "인증에 실패했습니다.",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockService := new(MockAuthService)
				tt.mockFn(mockService)
				handler := api.NewAuthHandler(mockService)

				w := httptest.NewRecorder()
				_, r := gin.CreateTestContext(w)
				r.POST("/login", handler.Login)

				jsonInput, _ := json.Marshal(tt.input)
				req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
				req.Header.Set("Content-Type", "application/json")

				r.ServeHTTP(w, req)

				assert.Equal(t, tt.expectedStatus, w.Code)
				assert.Contains(t, w.Body.String(), tt.expectedBody)
				mockService.AssertExpectations(t)
			})
		}
	})
}
