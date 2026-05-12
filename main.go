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

	// User
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	// Board Member
	boardMemberRepository := repository.NewBoardMemberRepository()

	// Board
	boardRepository := repository.NewBoardRepository()
	boardService := service.NewBoardService(boardRepository, userRepository, boardMemberRepository)
	boardController := controller.NewBoardController(boardService)

	// List Position
	listPositionRepository := repository.NewListPositionRepository()

	// List
	listRepository := repository.NewListRepository()
	listService := service.NewListService(listRepository, boardRepository, listPositionRepository)
	listController := controller.NewListController(listService)

	routes.Setup(app, userController, boardController, listController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port: ", port)

	log.Fatal(app.Listen(":" + port))
}
