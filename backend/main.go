package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MishraShardendu22/database"
	"github.com/MishraShardendu22/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var collections *mongo.Collection

func main() {
	fmt.Println("This is a Blog application")

	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Start Server
	StartServer(app)

	// // // // TESTING PLESASE IGNORE // // // //

	// recipient := "shardendumishra02@gmail.com"
	// otp := 123456
	// utils.MailSender(recipient, otp)

	// // // // TESTING PLESASE IGNORE // // // //

	// Listening To CORS
	SettingUpCORS(app)

	// Connecting To Database
	collections = database.Connect()
	SetUpRoutes(app, collections)

	port := os.Getenv("PORT")
	fmt.Println("Listening to port: " + port)
	if err := app.Listen("0.0.0.0:" + port); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func SetUpRoutes(app *fiber.App, collections *mongo.Collection) {
	// Signup Route
	routes.Signup(app, collections)

	// Login Route
	routes.Login(app, collections)

	// Like Route
	routes.Like(app, collections)

	// Blog Route
	routes.Blog(app, collections)

	// Comment Route
	routes.Comment(app, collections)
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
