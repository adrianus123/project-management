package routes

import (
	"log"

	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/controller"
	"github.com/adrianus123/project-management/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App, uc *controller.UserController, bc *controller.BoardController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)

	api := app.Group("/api/v1", middleware.JWTMiddleware(config.AppConfig.JWTSecret))

	userGroup := api.Group("/users")
	userGroup.Get("", uc.GetUserPagination)
	userGroup.Get("/:id", uc.GetUser)
	userGroup.Put("/:id", uc.UpdateUser)
	userGroup.Delete("/:id", uc.DeleteUser)

	boardGroup := api.Group("/boards")
	boardGroup.Post("", bc.CreateBoard)
	boardGroup.Put("/:id", bc.UpdateBoard)
	boardGroup.Post("/:id/members", bc.AddBoardMembers)
}
