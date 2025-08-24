package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	api := h.router.Group("api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

type User struct {
	ID   int
	Name string
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	names := []string{"Tim", "Bob"}
	users := []User{
		{ID: 1, Name: "Tim"},
		{ID: 2, Name: "Bob"},
	}

	data := struct {
		Names []string
		Users []User
	}{Names: names, Users: users}

	return c.Render("page", data)

	// return c.Render("page", fiber.Map{
	// 	"Count": 100,
	// 	"IsAdmin": false,
	// })
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	log.Info("Info")
	return c.SendString("Error!")
}
