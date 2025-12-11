package repositories_test

import (
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Auto migrate
	db.AutoMigrate(&models.User{})

	return db
}

func TestUserRepository_CreateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "hashedpassword",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1234567890",
		Gender:    models.Male,
	}

	err := repo.CreateUser(user)

	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.NotZero(t, user.CreatedAt)
}

func TestUserRepository_FindAllUsers(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewUserRepository(db)

	// Create test users
	user1 := &models.User{
		Name:      "User 1",
		Email:     "user1@example.com",
		Password:  "password1",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1111111111",
		Gender:    models.Male,
	}
	user2 := &models.User{
		Name:      "User 2",
		Email:     "user2@example.com",
		Password:  "password2",
		Birthdate: time.Date(1991, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "2222222222",
		Gender:    models.Female,
	}

	repo.CreateUser(user1)
	repo.CreateUser(user2)

	users, err := repo.FindAllUsers()

	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestUserRepository_FindUserByID(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "hashedpassword",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1234567890",
		Gender:    models.Male,
	}

	repo.CreateUser(user)

	var foundUser models.User
	foundUser, err := repo.FindUserByID(foundUser, int(user.ID))

	assert.NoError(t, err)
	assert.Equal(t, user.Name, foundUser.Name)
	assert.Equal(t, user.Email, foundUser.Email)
}

func TestUserRepository_FindUserByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewUserRepository(db)

	var user models.User
	_, err := repo.FindUserByID(user, 999)

	assert.Error(t, err)
}

func TestUserRepository_UpdateUser(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "hashedpassword",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1234567890",
		Gender:    models.Male,
	}

	repo.CreateUser(user)

	user.Name = "Updated Name"
	user.Email = "updated@example.com"

	err := repo.UpdateUser(user)

	assert.NoError(t, err)

	var updatedUser models.User
	updatedUser, _ = repo.FindUserByID(updatedUser, int(user.ID))
	assert.Equal(t, "Updated Name", updatedUser.Name)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}

func TestUserRepository_DeleteUser(t *testing.T) {
	db := setupTestDB(t)
	repo := repositories.NewUserRepository(db)

	user := &models.User{
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "hashedpassword",
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:     "1234567890",
		Gender:    models.Male,
	}

	repo.CreateUser(user)

	err := repo.DeleteUser(user.ID)

	assert.NoError(t, err)

	// Verify soft delete
	var foundUser models.User
	_, err = repo.FindUserByID(foundUser, int(user.ID))
	assert.Error(t, err)
}
