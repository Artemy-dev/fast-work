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

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.Render("page", fiber.Map{
		"Count": 1,
	})
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	log.Info("Info")
	return c.SendString("Error!")
}
