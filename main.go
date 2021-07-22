package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"wiatt/api"
)

//go:embed web/dist
var efs embed.FS

func main() {
	fx.New(
		fx.Options(api.Module),
		fx.Invoke(start),
	).Run()
}

func start(app *fiber.App, log *zap.SugaredLogger) {
	app.Get("/api/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello from Wiatt")
	})

	web, err := fs.Sub(efs, "web/dist")
	if err != nil {
		log.Fatalw("Impossible to properly select the embedded web folder !", "error", err)
	}

	app.Use(filesystem.New(filesystem.Config{
		Root:         http.FS(web),
		NotFoundFile: "index.html",
	}))
}
