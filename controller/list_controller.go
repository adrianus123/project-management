package controller

import (
	"github.com/adrianus123/project-management/constant"
	"github.com/adrianus123/project-management/model"
	"github.com/adrianus123/project-management/service"
	"github.com/adrianus123/project-management/util"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ListController struct {
	listService service.ListService
}

func NewListController(listService service.ListService) *ListController {
	return &ListController{
		listService: listService,
	}
}

func (c *ListController) CreateList(ctx fiber.Ctx) error {
	list := new(model.List)
	if err := ctx.Bind().Body(list); err != nil {
		return util.BadRequest(ctx, constant.ERR_BAD_REQUEST, nil, err.Error())
	}

	if err := c.listService.Create(list); err != nil {
		return util.InternalServerError(ctx, constant.ERR_INTERNAL_SERVER_ERROR, nil, err.Error())
	}

	return util.Success(ctx, constant.SUCCESS_CREATE_DATA, list)
}

func (c *ListController) UpdateList(ctx fiber.Ctx) error {
	publicID := ctx.Params("id")
	list := new(model.List)

	if err := ctx.Bind().Body(&list); err != nil {
		return util.BadRequest(ctx, constant.ERR_BAD_REQUEST, nil, err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return util.BadRequest(ctx, constant.ERR_BAD_REQUEST, nil, err.Error())
	}

	existingList, err := c.listService.GetByPublicID(publicID)
	if err != nil {
		return util.NotFound(ctx, constant.ERR_DATA_NOT_FOUND, nil, err.Error())
	}

	list.InternalID = existingList.InternalID
	list.PublicID = existingList.PublicID

	if err := c.listService.Update(list); err != nil {
		return util.InternalServerError(ctx, constant.ERR_INTERNAL_SERVER_ERROR, nil, err.Error())
	}

	updatedList, err := c.listService.GetByPublicID(publicID)
	if err != nil {
		return util.NotFound(ctx, constant.ERR_DATA_NOT_FOUND, nil, err.Error())
	}

	return util.Success(ctx, constant.SUCCESS_UPDATE_DATA, updatedList)
}

func (c *ListController) GetListOnBoard(ctx fiber.Ctx) error {
	boardPublicID := ctx.Params("board_id")

	if _, err := uuid.Parse(boardPublicID); err != nil {
		return util.BadRequest(ctx, constant.ERR_BAD_REQUEST, nil, err.Error())
	}

	lists, err := c.listService.GetByBoardID(boardPublicID)
	if err != nil {
		return util.NotFound(ctx, constant.ERR_DATA_NOT_FOUND, nil, err.Error())
	}

	return util.Success(ctx, constant.SUCCESS_GET_DATA, lists)
}
