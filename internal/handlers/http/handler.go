package http

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/usename-Poezd/zeus-instagram-dm/docs"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) Init(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	
	app.Get("/ping", h.Ping)
}

// Ping
// @Summary Ping
// @Tags service
// @Description Ping
// @ModuleID Зштп
// @Accept  json
// @Produce  json
// @Success 200 {string} string "pong"
// @Failure 400,401,500,503 {string} string "error"
// @Router /ping [get]
func (h Handler) Ping (c *fiber.Ctx) error {
	return c.Status(200).SendString("pong")
}