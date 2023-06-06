package controller

import (
	"dbKursach/internal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) UserRoute(route fiber.Router, mv *middlewares.MDWManager) {
	route.Post("/create_user", c.CreateUser())
	route.Post("/update_user_password", c.UpdateUserPassword())
	route.Post("/authorization", c.Authorization())
	route.Post("/withdraw", c.Withdraw())
	route.Get("/get_currencies", c.GetCurrencies())
	route.Post("/get_cards", c.GetCards())
	route.Post("/get_blocked_cards", c.GetBlockedCards())
	route.Post("/get_user_info", c.GetUserInfo())
	route.Post("/get_orders", c.GetOrders())
	route.Post("/get_orders_info", c.GetOrderInfo())
	route.Post("/get_card_orders", c.GetCardOrders())
	route.Post("/create_card", c.CreateCard())
}

func (c *AdmController) AdmRoute(route fiber.Router, mv *middlewares.MDWManager) {
	route.Post("/get_card_validation", c.AdmGetCardValidation())
	route.Post("/validate_card", c.AdmValidateCard())
	route.Post("/deny_validation_card", c.AdmDenyValidationCard())
	route.Post("/get_user_cards", c.AdmGetUserCards())
	route.Post("/get_user_card_info", c.AdmGetUserCardInfo())
	route.Post("/get_user_orders", c.AdmGetUserOrders())
	route.Post("/get_user_order_info", c.AdmGetUserOrderInfo())
	route.Post("/block_user_card", c.AdmBlockUserCard())
	route.Post("/unblock_user_card", c.AdmUnblockUserCard())
}
