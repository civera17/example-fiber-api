package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/civera17/fintech-test/database"
	"github.com/civera17/fintech-test/routes"
)

func main() {
	database.ConnectDb()
	app := SetupAPI()

	log.Fatal(app.Listen(":3000"))
}

func setUpRoutes(app *fiber.App) {
	// Endpoint that responds with slowest queries sorted by time and type and with pagination
	app.Get("/slowest-queries/:page/size/:size/type/:type", routes.SlowestQueries)
	app.Get("/allbooks", routes.AllBooks)
	app.Post("/addbook", routes.AddBook)
	app.Put("/update", routes.Update)
	app.Delete("/delete", routes.Delete)
}

func SetupAPI() *fiber.App {
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return app
}