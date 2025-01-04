package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MishraShardendu22/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("This is a Blog application")

	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	StartServer(app)

	// Listening To CORS
	SettingUpCORS(app)

	// Connecting To Database
	database.Connect()

	

	port := os.Getenv("PORT")
	fmt.Println("Listening to port: " + port)
	if err := app.Listen("0.0.0.0:" + port); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func SettingUpCORS(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE",
	}))
}

func StartServer(app *fiber.App) {
	fmt.Println("Server is starting")
	app.Get("/", SayHello)
}

func SayHello(c *fiber.Ctx) error {
	return c.SendString("Server Started !!")
}
