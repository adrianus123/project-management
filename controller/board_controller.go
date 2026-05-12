package controller

import (
	"github.com/adrianus123/project-management/constant"
	"github.com/adrianus123/project-management/model"
	"github.com/adrianus123/project-management/service"
	"github.com/adrianus123/project-management/util"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BoardController struct {
	boardService service.BoardService
}

func NewBoardController(boardService service.BoardService) *BoardController {
	return &BoardController{boardService: boardService}
}

func (c *BoardController) CreateBoard(ctx fiber.Ctx) error {
	var board model.Board

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID, err := uuid.Parse(claims["pub_id"].(string))
	if err != nil {
		return util.BadRequest(ctx, "", nil, err.Error())
	}

	board.OwnerPublicID = userID

	if err := ctx.Bind().Body(&board); err != nil {
		return util.BadRequest(ctx, "", nil, err.Error())
	}

	if err := c.boardService.Create(&board); err != nil {
		return util.BadRequest(ctx, "", nil, err.Error())
	}

	return util.Success(ctx, constant.SUCESS_CREATE, board)
}

func (c *BoardController) UpdateBoard(ctx fiber.Ctx) error {
	publicID := ctx.Params("id")
	board := new(model.Board)

	if err := ctx.Bind().Body(&board); err != nil {
		return util.BadRequest(ctx, constant.ERR_FAILED_PARSE_DATA, nil, err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return util.BadRequest(ctx, constant.ERR_INVALID_REQUEST, nil, err.Error())
	}

	existingBoard, err := c.boardService.GetByPublicID(publicID)
	if err != nil {
		return util.NotFound(ctx, constant.ERR_DATA_NOT_FOUND, nil, err.Error())
	}

	board.InternalID = existingBoard.InternalID
	board.PublicID = existingBoard.PublicID
	board.OwnerID = existingBoard.OwnerID
	board.OwnerPublicID = existingBoard.OwnerPublicID
	board.CreatedAt = existingBoard.CreatedAt

	if err := c.boardService.Update(board); err != nil {
		return util.InternalServerError(ctx, constant.ERR_FAILED_UPDATE_DATA, nil, err.Error())
	}

	return util.Success(ctx, constant.SUCCESS_UPDATE_DATA, board)
}
