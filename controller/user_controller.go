package controller

import (
	"github.com/adrianus123/project-management/model"
	"github.com/adrianus123/project-management/service"
	"github.com/adrianus123/project-management/util"
	"github.com/gofiber/fiber/v3"
	"github.com/jinzhu/copier"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx fiber.Ctx) error {
	user := new(model.User)

	if err := ctx.Bind().Body(user); err != nil {
		return util.BadRequest(ctx, "Failed parsing data", nil, err.Error())
	}

	if err := c.userService.Register(user); err != nil {
		return util.BadRequest(ctx, "Register failed", nil, err.Error())
	}

	var response model.UserResponse
	err := copier.Copy(&response, user)
	if err != nil {
		return util.InternalServerError(ctx, "Failed construct response", nil, err.Error())
	}

	return util.Success(ctx, "Register success", response)
}
