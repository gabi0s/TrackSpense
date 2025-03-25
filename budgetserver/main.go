package main

import (
	"budgetserver/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialise la base de données via GORM
	services.InitDB()

	// Initialise Fiber
	app := fiber.New()

	// Applique le middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type",
	}))

	// Routes Budget
	app.Get("/get-budget/:userid", services.GetBudgetsByUserID)
	app.Post("/create-budget", services.CreateBudget)
	app.Put("/update-budget/:userid/:budgetsid", services.UpdateBudget)
	app.Delete("/delete-budget/:userid/:budgetsid", services.DeleteBudget)

	port := os.Getenv("SERV_PORT_BUDGET")
	log.Println("Budget service launched on port :", port)

	// Démarre le serveur
	log.Fatal(app.Listen(":" + port))
}
