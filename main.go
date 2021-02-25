package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zamiramir/otn-gofiber-pizza/database"
	"github.com/zamiramir/otn-gofiber-pizza/pizza"
)

func setupRoutes(app *fiber.App) {
	app.Static("/", "./fiber-app/dist")

	app.Get("/api/v1/pizza", pizza.GetPizzas)
	app.Get("/api/v1/pizza/:id", pizza.GetPizza)
	app.Post("/api/v1/pizza", pizza.NewPizza)
	app.Delete("/api/v1/pizza/:id", pizza.DeletePizza)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "pizzas.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&pizza.Pizza{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	app.Listen(":3000")

	defer database.DBConn.Close()
}
