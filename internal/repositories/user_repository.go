package repositories

import (
	"socialVoleiAPI/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Where("deleted_at IS NULL").Find(&users).Error
	return users, err
}

func (r *UserRepository) FindUserByID(user models.User, intId int) (models.User, error) {
	err := r.db.Where("deleted_at IS NULL").First(&user, intId).Error
	return user, err
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Where("deleted_at IS NULL").Save(user).Error
}

func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Where("deleted_at IS NULL").Delete(&models.User{}, id).Error
}
