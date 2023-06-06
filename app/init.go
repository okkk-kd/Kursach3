package app

import (
	"context"
	ctrl "dbKursach/internal/controller"
	"dbKursach/internal/middlewares"
	"dbKursach/internal/repository"
	"dbKursach/internal/usecase"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *App) MapHandlers(ctx context.Context, app *fiber.App) error {
	rand.Seed(time.Now().UnixNano())

	middleware := middlewares.NewMDWManager(s.cfg, s.logger, s.pgDB)
	//Init repositories
	repository := repository.NewRepository(s.pgDB)

	// Init UseCases
	usecase := usecase.NewUsecase(repository)

	// Init delivery
	controller := ctrl.NewController(usecase, s.logger, s.cfg)
	admController := ctrl.NewAdmController(usecase, s.logger, s.cfg)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	apiGroup := app.Group("api")
	adminGroup := apiGroup.Group("control")

	controller.UserRoute(apiGroup, &middleware)
	admController.AdmRoute(adminGroup, &middleware)
	log.Default().Print("Handlers inited")
	return nil
}
