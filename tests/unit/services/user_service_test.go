package services_test

import (
	"errors"
	"socialVoleiAPI/internal/dto"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository é um mock do UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindAllUsers() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) FindUserByID(user models.User, intId int) (models.User, error) {
	args := m.Called(user, intId)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	req := &dto.CreateUserRequest{
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "password123",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1234567890",
		Gender:    models.Male,
	}

	expectedUser := &models.User{
		Name:      req.Name,
		Email:     req.Email,
		Birthdate: req.Birthdate,
		Phone:     req.Phone,
		Gender:    req.Gender,
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(0).(*models.User)
		user.ID = 1
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
	})

	user, err := service.CreateUser(req)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.NotEmpty(t, user.Password) // Password should be hashed
	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateUser_ValidationError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	req := &dto.CreateUserRequest{
		Name: "", // Empty name should fail validation
	}

	_, err := service.CreateUser(req)

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "CreateUser")
}

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	expectedUsers := []models.User{
		{BaseModel: models.BaseModel{ID: 1}, Name: "User 1", Email: "user1@example.com"},
		{BaseModel: models.BaseModel{ID: 2}, Name: "User 2", Email: "user2@example.com"},
	}

	mockRepo.On("FindAllUsers").Return(expectedUsers, nil)

	users, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, expectedUsers[0].Name, users[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	expectedUser := models.User{
		BaseModel: models.BaseModel{ID: 1},
		Name:      "Test User",
		Email:     "test@example.com",
	}

	mockRepo.On("FindUserByID", mock.Anything, 1).Return(expectedUser, nil)

	user, err := service.GetUserByID("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Email, user.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID_InvalidID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	_, err := service.GetUserByID("invalid")

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "FindUserByID")
}

func TestUserService_UpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	existingUser := models.User{
		BaseModel: models.BaseModel{ID: 1},
		Name:      "Old Name",
		Email:     "old@example.com",
		Phone:     "1111111111",
		Gender:    models.Male,
	}

	req := &dto.UpdateUserRequest{
		Name:  stringPtr("New Name"),
		Email: stringPtr("new@example.com"),
	}

	mockRepo.On("FindUserByID", mock.Anything, 1).Return(existingUser, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)

	user, err := service.UpdateUser("1", req)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "New Name", user.Name)
	assert.Equal(t, "new@example.com", user.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	req := &dto.UpdateUserRequest{
		Name: stringPtr("New Name"),
	}

	mockRepo.On("FindUserByID", mock.Anything, 1).Return(models.User{}, errors.New("record not found"))

	_, err := service.UpdateUser("1", req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "usuário não encontrado")
	mockRepo.AssertNotCalled(t, "UpdateUser")
}

func TestUserService_PutUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	existingUser := models.User{
		BaseModel: models.BaseModel{ID: 1},
		Name:      "Old Name",
		Email:     "old@example.com",
	}

	req := &dto.CreateUserRequest{
		Name:      "New Name",
		Email:     "new@example.com",
		Password:  "newpassword",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1234567890",
		Gender:    models.Female,
	}

	mockRepo.On("FindUserByID", mock.Anything, 1).Return(existingUser, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("*models.User")).Return(nil)

	user, err := service.PutUser("1", req)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, req.Name, user.Name)
	assert.Equal(t, req.Email, user.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	existingUser := models.User{
		BaseModel: models.BaseModel{ID: 1},
		Name:      "Test User",
	}

	mockRepo.On("FindUserByID", mock.Anything, 1).Return(existingUser, nil)
	mockRepo.On("DeleteUser", uint(1)).Return(nil)

	err := service.DeleteUser("1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := services.NewUserServiceWithRepository(mockRepo)

	mockRepo.On("FindUserByID", mock.Anything, 1).Return(models.User{}, errors.New("record not found"))

	err := service.DeleteUser("1")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "usuário não encontrado")
	mockRepo.AssertNotCalled(t, "DeleteUser")
}

// Helper function
func stringPtr(s string) *string {
	return &s
}
