package routes

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {
	// Inicializa camadas

	main := router.Group("/api/v1")

	{
		UserRoutes(main)
		// AuthRoutes(main)
		// BooksRoutes(main)
		// UserRoutes(main)
		// LoginRoutes(main)
	}

	return router
}
