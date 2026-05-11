package util

import "github.com/gofiber/fiber/v3"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error,omitempty"`
}

type PaginationResponse struct {
	Status       string         `json:"status"`
	ResponseCode int            `json:"response_code"`
	Message      string         `json:"message,omitempty"`
	Data         interface{}    `json:"data,omitempty"`
	Error        string         `json:"error,omitempty"`
	Meta         PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page      int    `json:"page" example:"1"`
	Limit     int    `json:"limit" example:"10"`
	Total     int    `json:"total" example:"100"`
	TotalPage int    `json:"total_page" example:"10"`
	Filter    string `json:"filter" example:"nama=Budi"`
	Sort      string `json:"sort" example:"-id"`
}

func BuildResponse(status string, responseCode int, message string, data interface{}, err string) Response {
	response := Response{
		Status:       status,
		ResponseCode: responseCode,
		Message:      message,
		Data:         data,
		Error:        "",
	}

	if err != "" {
		response.Error = err
	}

	return response
}

func BuildResponsePagination(status string, responseCode int, message string, data interface{}, meta PaginationMeta, err string) PaginationResponse {
	response := PaginationResponse{
		Status:       status,
		ResponseCode: responseCode,
		Message:      message,
		Data:         data,
		Error:        "",
		Meta:         meta,
	}

	if err != "" {
		response.Error = err
	}

	return response
}

func Success(c fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(BuildResponse("Success", fiber.StatusOK, message, data, ""))
}

func SuccessPagination(c fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(BuildResponsePagination("Success", fiber.StatusOK, message, data, meta, ""))
}

func Created(c fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(BuildResponse("Created", fiber.StatusCreated, message, data, ""))
}

func BadRequest(c fiber.Ctx, message string, data interface{}, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(BuildResponse("Bad Request", fiber.StatusBadRequest, message, data, err))
}

func NotFound(c fiber.Ctx, message string, data interface{}, err string) error {
	return c.Status(fiber.StatusNotFound).JSON(BuildResponse("Not Found", fiber.StatusNotFound, message, data, err))
}

func NotFoundPagination(c fiber.Ctx, message string, data interface{}, meta PaginationMeta, err string) error {
	return c.Status(fiber.StatusNotFound).JSON(BuildResponsePagination("Not Found", fiber.StatusNotFound, message, data, meta, err))
}

func InternalServerError(c fiber.Ctx, message string, data interface{}, err string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(BuildResponse("Internal Server Error", fiber.StatusInternalServerError, message, data, err))
}

func Unauthorized(c fiber.Ctx, message string, data interface{}, err string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(BuildResponse("Unauthorized", fiber.StatusUnauthorized, message, data, err))
}
