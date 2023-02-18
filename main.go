package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

const (
	ListenEnv            = "LISTEN"
	BasicAuthUserEnv     = "BASIC_AUTH_USER"
	BasicAuthPasswordEnv = "BASIC_AUTH_PASSWORD"
)

//go:embed site/*
var files embed.FS

func main() {
	var err error

	app := fiber.New(fiber.Config{})
	app.Use(logger.New())

	serveDir, err := fs.Sub(files, "site")
	if err != nil {
		log.Panic(err)
	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(serveDir),
		PathPrefix: "",
		Browse:     true,
	}))

	app.Use("/stats", basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv(BasicAuthUserEnv): os.Getenv(BasicAuthPasswordEnv),
		},
	}))
	app.Use("/stats", monitor.New(monitor.ConfigDefault))

	err = app.Listen(os.Getenv("LISTEN"))
	if err != nil {
		log.Panic(err)
	}

}
