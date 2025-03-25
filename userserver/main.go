package main

import (
	"log"
	"os"
	"userserver/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	services.InitDB() // inits db

	// inits Fiber
	app := fiber.New()

	// applies cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type",
	}))

	// GET routes
	app.Get("/get-solde/:userid", services.GetSoldeUser) // getting solde route

	// POST routes
	app.Post("/auth", services.AuthenticateUser)        // auth post route
	app.Post("/update-solde", services.UpdateSoldeUser) // updating solde route (adding or removing money for a user solde)

	// launchs the server
	port := os.Getenv("SERV_PORT_USER")
	log.Println("Server launched on port :  " + port)
	log.Fatal(app.Listen(":" + port))
}
