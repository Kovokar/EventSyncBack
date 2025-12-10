// controllers/user_controller.go
package controllers

import (
	"net/http"
	"socialVoleiAPI/internal/models"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.service.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, userToResponse(&user))
}

func (uc *UserController) GetUsers(ctx *gin.Context) {

	users, err := uc.service.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	userResponses := make([]models.UserResponse, len(users))

	for idx, user := range users {
		userResponses[idx] = userToResponse(&user)
	}

	ctx.JSON(http.StatusAccepted, userResponses)
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {

	user, err := uc.service.GetUserByID(ctx.Param("id"))

	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error: ": "Id Não Encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, userToResponse(&user))
}

func userToResponse(u *models.User) models.UserResponse {
	return models.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Birthdate: u.Birthdate,
		Email:     u.Email,
		Phone:     u.Phone,
		Gender:    u.Gender,
		Latitude:  u.Latitude,
		Longitude: u.Longitude,
		Photo:     u.Phone,
	}
}
