// controllers/user_controller.go
package controllers

import (
	"net/http"
	"socialVoleiAPI/internal/dto"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

// IUserService define a interface para o serviço de usuário
type IUserService interface {
	CreateUser(req *dto.CreateUserRequest) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	UpdateUser(id string, req *dto.UpdateUserRequest) (*models.User, error)
	PutUser(id string, req *dto.CreateUserRequest) (*models.User, error)
	DeleteUser(id string) error
}

type UserController struct {
	service IUserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

// NewUserControllerWithService permite criar um controller com qualquer implementação de IUserService
// Útil para testes
func NewUserControllerWithService(s IUserService) *UserController {
	return &UserController{service: s}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var req dto.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.service.CreateUser(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, userToResponse(user))
}

func (uc *UserController) GetUsers(ctx *gin.Context) {

	users, err := uc.service.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	userResponses := make([]dto.UserResponse, len(users))

	for idx, user := range users {
		userResponses[idx] = userToResponse(&user)
	}

	ctx.JSON(http.StatusAccepted, userResponses)
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {

	user, err := uc.service.GetUserByID(ctx.Param("id"))

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error: ": "Id Não Encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, userToResponse(&user))
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var req dto.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.service.UpdateUser(ctx.Param("id"), &req)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userToResponse(user))
}

func (uc *UserController) PutUser(ctx *gin.Context) {
	var req dto.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.service.PutUser(ctx.Param("id"), &req)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userToResponse(user))
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	err := uc.service.DeleteUser(ctx.Param("id"))
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func userToResponse(u *models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:                  u.BaseModel.ID,
		Name:                u.Name,
		Birthdate:           u.Birthdate,
		Email:               u.Email,
		Phone:               u.Phone,
		Gender:              u.Gender,
		Photo:               u.Photo,
		VisibleInPublicList: u.VisibleInPublicList,
		CreatedAt:           u.BaseModel.CreatedAt,
		UpdatedAt:           u.BaseModel.UpdatedAt,
	}
}
