package main

import (
	"TemplFiberHTMX/config"
	"TemplFiberHTMX/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()

	home.NewHandler(app)

	app.Listen(":3000")
}
