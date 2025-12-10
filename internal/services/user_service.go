package services

import (
	"fmt"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/repositories"
	"socialVoleiAPI/internal/utils"
	"socialVoleiAPI/internal/utils/validations"
	"strconv"
	"time"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *models.User) error {

	fmt.Println("res: ", req.Birthdate)

	if err := validations.ValidateRequiredFields(
		validations.Field{Name: "Nome", Value: req.Name},
		validations.Field{Name: "Senha", Value: req.Password},
		validations.Field{Name: "Email", Value: req.Email},
		validations.Field{Name: "Birthdate", Value: req.Birthdate},
		validations.Field{Name: "Latitude", Value: req.Latitude},
		validations.Field{Name: "Longitude", Value: req.Latitude},
	); err != nil {
		return err
	}

	parsedBirthdate, err := time.Parse("2006-01-02", req.Birthdate.Format("2006-01-02"))
	if err != nil {
		return fmt.Errorf("erro ao processar a data de nascimento: %v", err)
	}

	passwordHashed := utils.SHA256Encoder(req.Password)

	fmt.Println(parsedBirthdate)

	user := &models.User{
		Name:      req.Name,
		Birthdate: parsedBirthdate,
		Email:     req.Email,
		Password:  passwordHashed,
		Phone:     req.Phone,
		Gender:    models.GenderType(req.Gender),
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAllUsers()
}

func (s *UserService) GetUserByID(id string) (models.User, error) {

	var teste models.User
	intId, err := strconv.Atoi(id)
	if err != nil {
		return teste, fmt.Errorf("erro ao converter id para inteiro")
	}
	return s.repo.FindUserByID(teste, intId)
}
