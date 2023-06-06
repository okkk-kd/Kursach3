package controller

import (
	"dbKursach/internal/config"
	"dbKursach/internal/models"
	"dbKursach/internal/pkg/logger"
	"dbKursach/internal/pkg/utils/reqvalidator"
	"dbKursach/internal/usecase"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type AdmController struct {
	uc        usecase.Usecase
	loggingUC logger.LoggerUC
}

func NewAdmController(uc usecase.Usecase, log logger.Logger, cfg *config.Config) AdmController {
	return AdmController{
		uc:        uc,
		loggingUC: logger.NewLoggerUC(cfg, log),
	}
}

func (ctrl *AdmController) AdmGetCardValidation() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmGetCardsValidationReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmGetCardValidation(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmGetCardValidation()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *AdmController) AdmValidateCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmValidateCardReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmValidateCard(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmValidateCard()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *AdmController) AdmDenyValidationCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmDenyCardValidationReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmDenyValidationCard(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmDenyValidationCard()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *AdmController) AdmGetUserCards() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmGetUserCardsReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmGetUserCards(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmGetUserCards()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *AdmController) AdmGetUserCardInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmGetUserCardInfoReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmGetUserCardInfo(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmGetUserCardInfo()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *AdmController) AdmGetUserOrders() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmGetUserOrdersReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmGetUserOrders(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmGetUserOrders()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *AdmController) AdmGetUserOrderInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmGetUserOrderInfoReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmGetUserOrderInfo(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmGetUserOrderInfo()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *AdmController) AdmBlockUserCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmBlockUserCardReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmBlockUserCard(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmBlockUserCard()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *AdmController) AdmUnblockUserCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.AdmUnblockUserCardReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.AdmUnblockUserCard(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.AdmUnblockUserCard()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}
