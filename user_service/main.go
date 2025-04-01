package main

import (
	"go_microservice/user_service/config"
	"go_microservice/user_service/handlers"
	"go_microservice/user_service/repository"
	"go_microservice/user_service/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	db, err := config.InitDb()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	app := fiber.New()

	//initialize the logger
	logger := zap.Must(zap.NewProduction()).Sugar()

	// initialize the repositories
	repo := repository.NewRepository(db)

	// initialize the services
	svcs := services.NewServices(repo, logger)

	// initialize the handlers
	hndlrs := handlers.NewIndexHandler(svcs)

	// setup app routes
	config.InitServer(app, hndlrs)

	log.Panic(app.Listen(os.Getenv("SERVER_PORT")))
}
