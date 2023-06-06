package reqvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ReadRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(request); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), request)
}

func ValidateQueryParams(c *fiber.Ctx, request interface{}) error {
	if err := c.QueryParser(request); err != nil {
		return err
	}
	return validate.StructCtx(c.Context(), request)
}

func ValidateParams(c *fiber.Ctx, request interface{}) error {
	if err := c.ParamsParser(request); err != nil {
		return err
	}
	return validate.StructCtx(c.Context(), request)
}
