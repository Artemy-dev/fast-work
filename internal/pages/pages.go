package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}
	h.router.Get("/", h.index)
}

func (h *PagesHandler) index(c *fiber.Ctx) error {
	return c.SendString("Hello!")
}