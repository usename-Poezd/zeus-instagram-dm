package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/usename-Poezd/zeus-instagram-dm/internal/handlers/http"
)

// @title Zeus Instagram DM API
// @version 1.0
// @description REST API for Zeus Instagram DM spam

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

// Run initializes whole application.
func Run() {
	app := fiber.New()
	
	handler := http.NewHandler()
	handler.Init(app)

	app.Listen(":3000")
}