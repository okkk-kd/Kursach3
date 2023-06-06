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

type Controller struct {
	uc        usecase.Usecase
	loggingUC logger.LoggerUC
}

func NewController(uc usecase.Usecase, log logger.Logger, cfg *config.Config) Controller {
	return Controller{
		uc:        uc,
		loggingUC: logger.NewLoggerUC(cfg, log),
	}
}

func (ctrl *Controller) CreateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.CreateUser
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = ctrl.uc.CreateUser(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.CreateUser()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *Controller) UpdateUserPassword() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.UpdateUserPassword
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		err = ctrl.uc.UpdateUserPassword(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.UpdateUserPassword()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (ctrl *Controller) Authorization() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.Authorization
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.Authorization(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.Authorization()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) Withdraw() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.WithdrawReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.Withdraw(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.Withdraw()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) CreateCard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.CreateCardReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.CreateCard(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.CreateCard()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) GetCurrencies() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		result, err := ctrl.uc.GetCurrencies()
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.GetCurrencies()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) GetCards() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.GetCardsReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.GetCards(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.GetCards()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) GetBlockedCards() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.GetBlockedCardsReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.GetBlockedCards(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.GetBlockedCards()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) GetUserInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.GetUserInfoReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.GetUserInfo(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.GetUserInfo()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) GetOrders() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.GetOrdersReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.GetOrders(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.GetOrders()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) GetOrderInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.GetOrderInfoReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.GetOrderInfo(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.GetOrderInfo()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (ctrl *Controller) GetCardOrders() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params models.GetCardOrdersReq
		defer ctrl.loggingUC.CreateAPILog(ctx, time.Now())
		err := reqvalidator.ReadRequest(ctx, &params)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		result, err := ctrl.uc.GetCardOrders(params)
		if err != nil {
			err = errors.Wrapf(err, "userCtrl.GetCardOrders()")
			ctx.Locals("error", err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}
