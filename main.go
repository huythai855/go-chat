package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {

	// Create views engine
	viewsEngine := html.New("./views", ".html")

	// Start new fiber instance
	app := fiber.New(fiber.Config{
		Views: viewsEngine,
	})

	// Static route and directory
	app.Static("/static/", "./static")

	// Create a "ping" handler to test the server
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fiber")
	})

	// create new App Handler
	appHandler := NewAppHandler()

	// add appHandler routes
	app.Get("/", appHandler.HandleGetIndex)

	// create new websocket
	server := NewWebSocket()
	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessages()

	// Start the http server
	app.Listen(":3000")

}
