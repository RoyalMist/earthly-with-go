package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"wiatt/api/settings/config"
)

// Module makes the injectable available for FX.
var Module = fx.Provide(New)

// New creates a new injectable.
func New(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, config *config.Config) (app *fiber.App, err error) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			logger.Info(fmt.Sprintf("start API server started at %s", config.ApiHost()))
			go func() {
				err = app.Listen(config.ApiHost())
				err = fmt.Errorf("impossible to start API server at %s, %v", config.ApiHost(), err)
				logger.Fatal(err)
			}()

			return
		},
	})

	app = fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})

	return
}
