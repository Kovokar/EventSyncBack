package cmd

import (
	"socialVoleiAPI/internal/config"
	"socialVoleiAPI/internal/database"
	"socialVoleiAPI/internal/server"
)

func Run() {
	config.LoadEnvVariables()

	database.StartDb()

	server := server.NewServer()
	server.Run()
}
