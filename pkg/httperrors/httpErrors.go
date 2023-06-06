package httperrors

import (
	"dbKursach/internal/config"
	"dbKursach/internal/pkg/logger"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type responseMsg struct {
	Message string `json:"message"`
}

func Init(config *config.Config, logger logger.Logger) func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		var (
			response   responseMsg
			statusCode int
		)

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"app_user_username_uindex\"") {
			response.Message = "username unique constraint"
			statusCode = fiber.StatusBadRequest
		} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"app_user_phone_uindex\"") {
			response.Message = "phone unique constraint"
			statusCode = fiber.StatusBadRequest
		} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"app_user_email_uindex\"") {
			response.Message = "email unique constraint"
			statusCode = fiber.StatusBadRequest
		} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"qiwi_vault_name_uindex\"") {
			response.Message = "wallet with this name already exists"
			statusCode = fiber.StatusConflict
		} else if strings.Contains(err.Error(), "Field validation for") {
			response.Message = "Неверные данные"
			statusCode = fiber.StatusBadRequest
		} else if err.Error() == "not found confirmation code" {
			response.Message = err.Error()
			statusCode = fiber.StatusBadRequest
		} else if err.Error() == "wrong current password hash" {
			response.Message = err.Error()
			statusCode = fiber.StatusBadRequest
		} else if err.Error() == "session doesnt found" {
			response.Message = err.Error()
			statusCode = fiber.StatusUnauthorized
		}

		if statusCode == 0 {
			statusCode = fiber.StatusInternalServerError
		}

		return c.Status(statusCode).JSON(response)
	}
}
