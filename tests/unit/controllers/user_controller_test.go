package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"socialVoleiAPI/internal/controllers"
	"socialVoleiAPI/internal/dto"
	"socialVoleiAPI/internal/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Helper para criar controller com mock
func createControllerWithMock(mockService *MockUserService) *controllers.UserController {
	return controllers.NewUserControllerWithService(mockService)
}

// MockUserService é um mock do UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(req *dto.CreateUserRequest) (*models.User, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id string) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(id string, req *dto.UpdateUserRequest) (*models.User, error) {
	args := m.Called(id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) PutUser(id string, req *dto.CreateUserRequest) (*models.User, error) {
	args := m.Called(id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestUserController_CreateUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.POST("/users", controller.CreateUser)

	reqBody := dto.CreateUserRequest{
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "password123",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1234567890",
		Gender:    models.Male,
	}

	expectedUser := &models.User{
		BaseModel: models.BaseModel{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:      reqBody.Name,
		Email:     reqBody.Email,
		Birthdate: reqBody.Birthdate,
		Phone:     reqBody.Phone,
		Gender:    reqBody.Gender,
	}

	mockService.On("CreateUser", mock.AnythingOfType("*dto.CreateUserRequest")).Return(expectedUser, nil)

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestUserController_CreateUser_InvalidRequest(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.POST("/users", controller.CreateUser)

	reqBody := map[string]interface{}{
		"name": "", // Invalid: empty name
	}

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "CreateUser")
}

func TestUserController_GetUsers(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.GET("/users", controller.GetUsers)

	expectedUsers := []models.User{
		{
			BaseModel: models.BaseModel{ID: 1},
			Name:      "User 1",
			Email:     "user1@example.com",
		},
		{
			BaseModel: models.BaseModel{ID: 2},
			Name:      "User 2",
			Email:     "user2@example.com",
		},
	}

	mockService.On("GetAllUsers").Return(expectedUsers, nil)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
	mockService.AssertExpectations(t)
}

func TestUserController_GetUserByID(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.GET("/users/:id", controller.GetUserByID)

	expectedUser := models.User{
		BaseModel: models.BaseModel{ID: 1},
		Name:      "Test User",
		Email:     "test@example.com",
	}

	mockService.On("GetUserByID", "1").Return(expectedUser, nil)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
	mockService.AssertExpectations(t)
}

func TestUserController_GetUserByID_NotFound(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.GET("/users/:id", controller.GetUserByID)

	mockService.On("GetUserByID", "999").Return(models.User{}, errors.New("record not found"))

	req, _ := http.NewRequest("GET", "/users/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestUserController_UpdateUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.PATCH("/users/:id", controller.UpdateUser)

	reqBody := dto.UpdateUserRequest{
		Name:  stringPtr("Updated Name"),
		Email: stringPtr("updated@example.com"),
	}

	expectedUser := &models.User{
		BaseModel: models.BaseModel{ID: 1},
		Name:      "Updated Name",
		Email:     "updated@example.com",
	}

	mockService.On("UpdateUser", "1", mock.AnythingOfType("*dto.UpdateUserRequest")).Return(expectedUser, nil)

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PATCH", "/users/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestUserController_UpdateUser_NotFound(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.PATCH("/users/:id", controller.UpdateUser)

	reqBody := dto.UpdateUserRequest{
		Name: stringPtr("Updated Name"),
	}

	mockService.On("UpdateUser", "999", mock.AnythingOfType("*dto.UpdateUserRequest")).Return(nil, errors.New("usuário não encontrado"))

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PATCH", "/users/999", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestUserController_PutUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.PUT("/users/:id", controller.PutUser)

	reqBody := dto.CreateUserRequest{
		Name:      "Updated User",
		Email:     "updated@example.com",
		Password:  "newpassword",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1234567890",
		Gender:    models.Female,
	}

	expectedUser := &models.User{
		BaseModel: models.BaseModel{ID: 1},
		Name:      reqBody.Name,
		Email:     reqBody.Email,
	}

	mockService.On("PutUser", "1", mock.AnythingOfType("*dto.CreateUserRequest")).Return(expectedUser, nil)

	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestUserController_DeleteUser(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.DELETE("/users/:id", controller.DeleteUser)

	mockService.On("DeleteUser", "1").Return(nil)

	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

func TestUserController_DeleteUser_NotFound(t *testing.T) {
	mockService := new(MockUserService)
	controller := createControllerWithMock(mockService)

	router := setupRouter()
	router.DELETE("/users/:id", controller.DeleteUser)

	mockService.On("DeleteUser", "999").Return(errors.New("usuário não encontrado"))

	req, _ := http.NewRequest("DELETE", "/users/999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

// Helper function
func stringPtr(s string) *string {
	return &s
}
