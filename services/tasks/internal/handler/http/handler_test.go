package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtemFed/hse-wishlist/pkg/xhttp"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/config"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/domain"
	"github.com/ArtemFed/hse-wishlist/services/tasks/internal/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockAuthService структура для мокирования AuthService
type MockAuthService struct {
	mock.Mock
}

// MockTaskService структура для мокирования TaskService
type MockTaskService struct {
	mock.Mock
}

// List метод для мокирования List
func (m *MockTaskService) List(ctx context.Context, filter domain.TaskFilter) ([]domain.TaskGet, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]domain.TaskGet), args.Error(1)
}

// Create метод для мокирования Create
func (m *MockTaskService) Create(ctx context.Context, task domain.TaskCreate) (*uuid.UUID, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(*uuid.UUID), args.Error(1)
}

// Update метод для мокирования Update
func (m *MockTaskService) Update(ctx context.Context, task domain.TaskUpdate) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

// Patch метод для мокирования Patch
func (m *MockTaskService) Patch(ctx context.Context, id uuid.UUID, status string) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

// MockAccountService структура для мокирования AccountService
type MockAccountService struct {
	mock.Mock
}

// Нет методов для AccountService в текущем примере, но можно добавить при необходимости

func TestHandler(t *testing.T) {
	// Инициализация логгера
	logger, _ := zap.NewDevelopment()
	log.Logger = logger

	// Создаем моки сервисов
	mockTaskService := new(MockTaskService)

	id := uuid.New()
	// Ожидаемые значения
	expectedTasks := []domain.TaskGet{
		{
			UUID: id,
			Name: "Task 1",
		},
	}

	expectedTaskUUID := id

	// Настраиваем моки
	mockTaskService.On("List", mock.Anything, domain.TaskFilter{}).Return(expectedTasks, nil)
	mockTaskService.On("Create", mock.Anything, domain.TaskCreate{Name: "New Task"}).Return(&expectedTaskUUID, nil)
	mockTaskService.On("Update", mock.Anything, domain.TaskUpdate{Name: "Updated Task"}).Return(nil)
	mockTaskService.On("Patch", mock.Anything, id, "Finish").Return(nil)

	// Создаем общий роутинг http сервера
	router := xhttp.NewRouter()

	mainHandler := NewHandler(
		&config.Config{},
		mockTaskService,
		nil,
		nil,
	)

	InitMainHandler(mainHandler, router.Router(), []MiddlewareFunc{}, "")

	// Создаем тестовый сервер
	ts := httptest.NewServer(router.Router())
	defer ts.Close()

	// Тестирование маршрута /api/v1/tasks с валидным токеном
	t.Run("Test GET /api/v1/tasks with valid token", func(t *testing.T) {
		url := fmt.Sprintf("%s/api/v1/tasks", ts.URL)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer valid_token")

		resp := httptest.NewRecorder()
		router.Router().ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var responseData []domain.TaskGet
		json.Unmarshal(resp.Body.Bytes(), &responseData)
		assert.Equal(t, expectedTasks, responseData)
	})

	// Тестирование маршрута /api/v1/tasks с POST запросом
	t.Run("Test POST /api/v1/tasks with valid token", func(t *testing.T) {
		url := fmt.Sprintf("%s/api/v1/tasks", ts.URL)
		requestBody := gin.H{
			"name": "New Task",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Authorization", "Bearer valid_token")
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.Router().ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var responseData uuid.UUID
		json.Unmarshal(resp.Body.Bytes(), &responseData)
		assert.Equal(t, id, responseData)
	})

	// Тестирование маршрута /api/v1/tasks с PUT запросом
	t.Run("Test PUT /api/v1/tasks with valid token", func(t *testing.T) {
		url := fmt.Sprintf("%s/api/v1/tasks", ts.URL)
		requestBody := gin.H{
			"name": "Updated Task",
		}
		requestBodyBytes, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
		req.Header.Set("Authorization", "Bearer valid_token")
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.Router().ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	// Тестирование маршрута /api/v1/tasks с PATCH запросом
	t.Run("Test Patch /api/v1/tasks with valid token", func(t *testing.T) {
		url := fmt.Sprintf("%s/api/v1/tasks?id=%s&status=%s", ts.URL, id, "Finish")

		req, _ := http.NewRequest("PATCH", url, nil)
		req.Header.Set("Authorization", "Bearer valid_token")
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.Router().ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
