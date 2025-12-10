package server

import (
	"log"
	routes "socialVoleiAPI/internal/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

func NewServer() Server {
	return Server{
		Port:   "8081",
		Server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigureRoutes(s.Server)
	log.Print("Server running on port :" + s.Port)
	log.Fatal(router.Run(":" + s.Port))
}
