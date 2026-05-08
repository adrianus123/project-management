package main

import (
	"log"

	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/controller"
	"github.com/adrianus123/project-management/database/seed"
	"github.com/adrianus123/project-management/repository"
	"github.com/adrianus123/project-management/routes"
	"github.com/adrianus123/project-management/service"
	"github.com/gofiber/fiber/v3"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)

	log.Fatal(app.Listen(":" + port))
}
