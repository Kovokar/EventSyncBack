package routes

import (
	"socialVoleiAPI/internal/controllers"
	"socialVoleiAPI/internal/database"
	"socialVoleiAPI/internal/repositories"
	"socialVoleiAPI/internal/services"

	"github.com/gin-gonic/gin"
)

// Função que registra as rotas relacionadas a "books"
func UserRoutes(router *gin.RouterGroup) {

	db := database.GetDatabase()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	user := router.Group("/user")

	{
		user.GET("/", userController.GetUsers)
		user.GET("/:id", userController.GetUserByID)
		user.POST("/", userController.CreateUser)

	}
}
