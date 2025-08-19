package main

import (
	"TemplFiberHTMX/config"
	"TemplFiberHTMX/internal/home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	engine := html.New("./html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New()) // 
	app.Use(recover.New())
	log.SetLevel(log.Level(logConfig.Level))
	home.NewHandler(app)

	app.Listen(":3000") // curl -v http://127.0.0.1:3000/
}

// git add .; git commit -m "