package main

import (
	services "expenseserver/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialise la base de données
	services.InitDB()

	// Initialise Fiber
	app := fiber.New()

	// Active CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type",
	}))

	// POST
	app.Post("/create-expense", services.AddExpense)

	// GET
	app.Get("/get-expense/:userid", services.GetExpense)

	// Lancement du serveur
	port := os.Getenv("SERV_PORT_EXPENSE")

	log.Println("Server launched on port: " + port)
	log.Fatal(app.Listen(":" + port))
}
