package controller

import (
	"math"
	"strconv"

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
	err := copier.Copy(&response, &user)
	if err != nil {
		return util.InternalServerError(ctx, "Failed construct response", nil, err.Error())
	}

	return util.Success(ctx, "Register success", response)
}

func (c *UserController) Login(ctx fiber.Ctx) error {
	var loginRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.Bind().Body(&loginRequest); err != nil {
		return util.BadRequest(ctx, "Invalid Request", nil, err.Error())
	}

	user, err := c.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return util.Unauthorized(ctx, "Invalid Credential", nil, err.Error())
	}

	token, _ := util.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := util.GenerateRefreshToken(user.InternalID)

	var data model.UserResponse
	err = copier.Copy(&data, &user)
	if err != nil {
		return util.InternalServerError(ctx, "Failed construct response", nil, err.Error())
	}

	return util.Success(ctx, "Login Success", fiber.Map{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          data,
	})
}

func (c *UserController) GetUser(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.userService.GetUserByPublicID(id)
	if err != nil {
		return util.NotFound(ctx, "User Not Found", nil, err.Error())
	}

	var response model.UserResponse
	err = copier.Copy(&response, &user)
	if err != nil {
		return util.InternalServerError(ctx, "Failed construct response", nil, err.Error())
	}

	return util.Success(ctx, "Get User Success", response)
}

func (c *UserController) GetUserPagination(ctx fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	offset := (page - 1) * limit

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	users, total, err := c.userService.GetAllPagination(filter, sort, limit, offset)
	if err != nil {
		return util.InternalServerError(ctx, "Failed get users", nil, err.Error())
	}

	var response []model.UserResponse
	if err := copier.Copy(&response, &users); err != nil {
		return util.InternalServerError(ctx, "Failed construct response", nil, err.Error())
	}

	meta := util.PaginationMeta{
		Page:      page,
		Limit:     limit,
		Total:     int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter:    filter,
		Sort:      sort,
	}

	if total == 0 {
		return util.NotFoundPagination(ctx, "", nil, meta, "No user found")
	}

	return util.SuccessPagination(ctx, "Get All Users Success", response, meta)
}
