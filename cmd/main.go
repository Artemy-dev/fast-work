package main

import (
	"TemplFiberHTMX/config"
	"TemplFiberHTMX/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	app.Listen(":3000")  // curl -v http://127.0.0.1:3000/
}
