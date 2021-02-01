package main

import (
	"github.com/JoWiel/component-set-generator/router"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	app := fiber.New()
	//Sets headers to improve security
	app.Use(helmet.New())
	//Handle Cors
	app.Use(cors.New())

	// database.ConnectDB()

	router.SetupRoutes(app)

	//Heroku automatically assigns a port our web server. If it   //fails we instruct it to use port 5000
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	log.Fatal(app.Listen(port))
	// defer database.DB.Close()
}
