package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed web/dist
var efs embed.FS

func main() {
	app := fiber.New()
	app.Get("/api/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello from Wiatt")
	})

	web, err := fs.Sub(efs, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	app.Use(filesystem.New(filesystem.Config{
		Root:         http.FS(web),
		NotFoundFile: "index.html",
	}))

	log.Fatal(app.Listen(":2000"))
}
