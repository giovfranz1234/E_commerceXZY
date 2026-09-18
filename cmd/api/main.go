// package main
package main

import (
	"os"

	"github.com/giovfranz1234/E_commerceXZY/internal/handler"
	"github.com/giovfranz1234/E_commerceXZY/internal/logger"
	"github.com/giovfranz1234/E_commerceXZY/internal/middlewares"
	"github.com/giovfranz1234/E_commerceXZY/internal/server"
)

func main() {

	logger := logger.New(logger.DefaultConfig())
	//Aqui se puede crear el servidor y configurarlo con las opciones deseadas
	srv := server.New(":8080", logger)

	srv.Use(middlewares.Recovery(logger))
	srv.Use(middlewares.Logger(logger))

	userHandler := handler.NewUserHandler()
	srv.RegisterRoutes("GET /listUser", userHandler)

	if err := srv.Start(); err != nil {
		logger.Error("error fatal", "error", err)
		os.Exit(1)
	}

}
