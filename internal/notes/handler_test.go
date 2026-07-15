package notes

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

type mockNotesService struct {
	mock.Mock
}

func (m *mockNotesService) Create(notes Note) (Note, error) {
	args := m.Called(notes)
	return args.Get(0).(Note), args.Error(1)
}
func (m *mockNotesService) GetAll() ([]Note, error) {
	args := m.Called()
	return args.Get(0).([]Note), args.Error(1)
}
func (m *mockNotesService) GetByID(id int) (*Note, error) {
	args := m.Called(id)
	return args.Get(0).(*Note), args.Error(1)
}
func (m *mockNotesService) Update(id int, note Note) (*Note, error) {
	args := m.Called(id)
	return args.Get(0).(*Note), args.Error(1)
}
func (m *mockNotesService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
func TestNotesHandlerCreate(t *testing.T) {
	t.Run("Create Notes successfully", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mockNotesService)
		handler := NewNoteHandler(mockService)
		app.Post("/notes", handler.Create)
		notes := Note{
			Title:   "Test Notes",
			Content: "Test Description",
		}
		expectedNotes := notes
		expectedNotes.ID = 1
		mockService.On("Create", notes).Return(expectedNotes, nil).Once()
		body, _ := json.Marshal(notes)
		req := httptest.NewRequest("Post", "/notes", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, 201, resp.StatusCode)
	})

	t.Run("Invalid json parser", func(t *testing.T) {

		app := fiber.New()
		mockservice := new(mockNotesService)
		handler := NewNoteHandler(mockservice)

		app.Post("/notes", handler.Create)
		invalidBody := `{"title": "Test Task", "description": invalid json}`
		req := httptest.NewRequest("POST", "/notes", bytes.NewReader([]byte(invalidBody)))
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		var errResponse map[string]any
		err = json.NewDecoder(resp.Body).Decode(&errResponse)
		assert.NoError(t, err)
		assert.Contains(t, errResponse, "error")

		mockservice.AssertNotCalled(t, "Create", mock.Anything)
	})

	t.Run("Service create error", func(t *testing.T) {
		app := fiber.New()
		mockService := new(mockNotesService)
		handler := NewNoteHandler(mockService)

		app.Post("/notes", handler.Create)

		notes := Note{
			Title:   "Test Notes",
			Content: "Test Description",
		}
		mockService.On("Create", notes).Return(Note{}, errors.New("database error")).Once()
		body, _ := json.Marshal(notes)
		req := httptest.NewRequest("Post", "/notes", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		var errorResponse map[string]any
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)

		assert.NoError(t, err)
		assert.Equal(t, "database error", errorResponse["error"])
		mockService.AssertExpectations(t)
	})
}
