package services

import (
	"fmt"
	"socialVoleiAPI/internal/dto"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/repositories"
	"socialVoleiAPI/internal/utils"
	"socialVoleiAPI/internal/utils/validations"
	"strconv"
	"time"
)

// IUserRepository define a interface para o repositório de usuário
type IUserRepository interface {
	CreateUser(user *models.User) error
	FindAllUsers() ([]models.User, error)
	FindUserByID(user models.User, intId int) (models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type UserService struct {
	repo IUserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// NewUserServiceWithRepository permite criar um service com qualquer implementação de IUserRepository
// Útil para testes
func NewUserServiceWithRepository(repo IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*models.User, error) {

	fmt.Println("res: ", req.Birthdate)

	if err := validations.ValidateRequiredFields(
		validations.Field{Name: "Nome", Value: req.Name},
		validations.Field{Name: "Senha", Value: req.Password},
		validations.Field{Name: "Email", Value: req.Email},
		validations.Field{Name: "Birthdate", Value: req.Birthdate},
	); err != nil {
		return nil, err
	}

	parsedBirthdate, err := time.Parse("2006-01-02", req.Birthdate.Format("2006-01-02"))
	if err != nil {
		return nil, fmt.Errorf("erro ao processar a data de nascimento: %v", err)
	}

	passwordHashed := utils.SHA256Encoder(req.Password)

	fmt.Println(parsedBirthdate)

	// Define valor padrão para Gender se não fornecido
	gender := req.Gender
	if gender == "" {
		gender = models.Other
	}

	user := &models.User{
		Name:                req.Name,
		Birthdate:           parsedBirthdate,
		Email:               req.Email,
		Password:            passwordHashed,
		Phone:               req.Phone,
		Gender:              gender,
		Photo:               req.Photo,
		VisibleInPublicList: req.VisibleInPublicList,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
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

func (s *UserService) UpdateUser(id string, req *dto.UpdateUserRequest) (*models.User, error) {
	var user models.User
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter id para inteiro")
	}

	// Busca o usuário existente
	user, err = s.repo.FindUserByID(user, intId)
	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	// Atualiza apenas os campos fornecidos
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		user.Password = utils.SHA256Encoder(*req.Password)
	}
	if req.Birthdate != nil {
		parsedBirthdate, err := time.Parse("2006-01-02", req.Birthdate.Format("2006-01-02"))
		if err != nil {
			return nil, fmt.Errorf("erro ao processar a data de nascimento: %v", err)
		}
		user.Birthdate = parsedBirthdate
	}
	if req.Phone != nil {
		user.Phone = *req.Phone
	}
	if req.Gender != nil {
		user.Gender = *req.Gender
	}
	if req.Photo != nil {
		user.Photo = *req.Photo
	}
	if req.VisibleInPublicList != nil {
		user.VisibleInPublicList = *req.VisibleInPublicList
	}

	if err := s.repo.UpdateUser(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) PutUser(id string, req *dto.CreateUserRequest) (*models.User, error) {
	var user models.User
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter id para inteiro")
	}

	// Busca o usuário existente
	user, err = s.repo.FindUserByID(user, intId)
	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	// Valida campos obrigatórios
	if err := validations.ValidateRequiredFields(
		validations.Field{Name: "Nome", Value: req.Name},
		validations.Field{Name: "Senha", Value: req.Password},
		validations.Field{Name: "Email", Value: req.Email},
		validations.Field{Name: "Birthdate", Value: req.Birthdate},
	); err != nil {
		return nil, err
	}

	// Atualiza todos os campos
	parsedBirthdate, err := time.Parse("2006-01-02", req.Birthdate.Format("2006-01-02"))
	if err != nil {
		return nil, fmt.Errorf("erro ao processar a data de nascimento: %v", err)
	}

	passwordHashed := utils.SHA256Encoder(req.Password)

	gender := req.Gender
	if gender == "" {
		gender = models.Other
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Password = passwordHashed
	user.Birthdate = parsedBirthdate
	user.Phone = req.Phone
	user.Gender = gender
	user.Photo = req.Photo
	user.VisibleInPublicList = req.VisibleInPublicList

	if err := s.repo.UpdateUser(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) DeleteUser(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("erro ao converter id para inteiro")
	}

	// Verifica se o usuário existe
	var user models.User
	user, err = s.repo.FindUserByID(user, intId)
	if err != nil {
		return fmt.Errorf("usuário não encontrado")
	}

	return s.repo.DeleteUser(uint(intId))
}
