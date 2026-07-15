package tasks

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTaskService is a mock implementation of TaskService
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) Create(task Task) (Task, error) {
	args := m.Called(task)
	return args.Get(0).(Task), args.Error(1)
}

func (m *MockTaskService) GetAll() ([]Task, error) {
	args := m.Called()
	return args.Get(0).([]Task), args.Error(1)
}

func (m *MockTaskService) GetByID(id int) (*Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Task), args.Error(1)
}

func (m *MockTaskService) Update(id int, task Task) (*Task, error) {
	args := m.Called(id, task)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Task), args.Error(1)
}

func (m *MockTaskService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestTaskHandler_Create(t *testing.T) {

	t.Run("successful task creation", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Post("/tasks", handler.Create)

		task := Task{
			Title:       "Test Task",
			Description: "Test Description",
			Completed:   false,
		}
		expectedTask := task
		expectedTask.ID = 1

		mockService.On("Create", task).Return(expectedTask, nil).Once()

		body, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)

		var responseTask Task
		err = json.NewDecoder(resp.Body).Decode(&responseTask)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask, responseTask)
		mockService.AssertExpectations(t)
	})

	t.Run("invalid JSON body", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Post("/tasks", handler.Create)

		invalidBody := `{"title": "Test Task", "description": invalid json}`
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader([]byte(invalidBody)))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		var errorResponse map[string]any
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse, "error")
		mockService.AssertNotCalled(t, "Create", mock.Anything)
	})

	t.Run("empty request body", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Post("/tasks", handler.Create)

		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader([]byte{}))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		var errorResponse map[string]any
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.NoError(t, err)
		assert.Contains(t, errorResponse, "error")
		mockService.AssertNotCalled(t, "Create", mock.Anything)
	})

	t.Run("service create error", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Post("/tasks", handler.Create)

		task := Task{
			Title:       "Test Task",
			Description: "Test Description",
			Completed:   false,
		}

		mockService.On("Create", task).Return(Task{}, errors.New("database error")).Once()

		body, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		var errorResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		assert.NoError(t, err)
		assert.Equal(t, "database error", errorResponse["error"])
		mockService.AssertExpectations(t)
	})

	t.Run("task with missing required fields", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Post("/tasks", handler.Create)

		task := Task{
			Description: "Test Description",
			Completed:   false,
		}

		mockService.On("Create", task).Return(task, nil).Once()

		body, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)

		var responseTask Task
		err = json.NewDecoder(resp.Body).Decode(&responseTask)
		assert.NoError(t, err)
		assert.Equal(t, task, responseTask)
		mockService.AssertExpectations(t)
	})

	t.Run("task with all fields populated", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Post("/tasks", handler.Create)

		task := Task{
			Title:       "Complete Task",
			Description: "A fully detailed task",
			Completed:   true,
		}
		expectedTask := task
		expectedTask.ID = 2

		mockService.On("Create", task).Return(expectedTask, nil).Once()

		body, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)

		var responseTask Task
		err = json.NewDecoder(resp.Body).Decode(&responseTask)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask, responseTask)
		mockService.AssertExpectations(t)
	})

	t.Run("malformed JSON with extra fields", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Post("/tasks", handler.Create)

		malformedBody := `{"title": "Test", "description": "Desc", "completed": false, "extra_field": "ignored"}`
		var task Task
		json.Unmarshal([]byte(malformedBody), &task) // This should work as extra fields are ignored

		mockService.On("Create", task).Return(task, nil).Once()

		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader([]byte(malformedBody)))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)
		mockService.AssertExpectations(t)
	})

	t.Run("large JSON payload", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Post("/tasks", handler.Create)

		task := Task{
			Title:       string(make([]byte, 1000)), // Large title
			Description: string(make([]byte, 5000)), // Large description
			Completed:   false,
		}

		mockService.On("Create", task).Return(task, nil).Once()

		body, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

func TestTaskHandler_GetAll(t *testing.T) {
	t.Run("calling all tasks successfully", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockTaskService)
		handler := NewTaskHandler(mockService)
		app.Get("/tasks", handler.GetAll)

		tasks := []Task{{
			ID:    1,
			Title: "Learn Go",
		}}
		mockService.On("GetAll").Return(tasks, nil).Once()
		req := httptest.NewRequest("GET", "/tasks", nil)
		resp, err := app.Test(req)
		// will check if everything is correct by passsing the err to this function
		assert.NoError(t, err)
		// will check if the response statuscode is same as expected
		assert.Equal(t, 200, resp.StatusCode)
		var response []Task
		json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, tasks, response)
		// Verify that the mock service methods were called as expected
		mockService.AssertExpectations(t)
	})

	t.Run("service failed during get all tasks", func(t *testing.T) {
		// app := fiber.New()
		// mockService := new(MockTaskService)
		// taskHandler := NewTaskHandler(mockService)
		// app.Get("/tasks", taskHandler.GetAll)

		// mockService.On("GetAll").Return([]Task{}, errors.New("database error")).Once()

		// req := httptest.NewRequest("GET", "/tasks", nil)
		// resp, err := app.Test(req)

		// assert.NoError(t, err)
		// assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		// var errorResponse map[string]any
		// err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		// assert.NoError(t, err)

		// assert.Equal(t, "database error", errorResponse["error"])

		// mockService.AssertExpectations(t)

		app := fiber.New()
		mockservice := new(MockTaskService)
		handler := NewTaskHandler(mockservice)
		app.Get("/tasks", handler.GetAll)
		mockservice.On("GetAll").Return([]Task{}, errors.New("database error")).Once()
		req := httptest.NewRequest("GET", "/tasks", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
		mockservice.AssertExpectations(t)

	})
}
