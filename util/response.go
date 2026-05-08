package util

import "github.com/gofiber/fiber/v3"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error,omitempty"`
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

func Success(c fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(BuildResponse("Success", fiber.StatusOK, message, data, ""))
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

func InternalServerError(c fiber.Ctx, message string, data interface{}, err string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(BuildResponse("Internal Server Error", fiber.StatusInternalServerError, message, data, err))
}
