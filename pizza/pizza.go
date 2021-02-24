package pizza

import (
	"fmt"

	"https://github.com/zamiramir/otn-gofiber-pizza/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Pizza struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetPizzas(c *fiber.Ctx) error {
	db := database.DBConn
	var pizzas []Pizza
	db.Find(&pizzas)
	return c.JSON(pizzas)
}

func GetPizza(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var pizza Pizza
	db.Find(&pizza, id)
	return c.JSON(pizza)
}

func NewPizza(c *fiber.Ctx) error {
	db := database.DBConn
	pizza := new(Pizza)
	if err := c.BodyParser(pizza); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	fmt.Println(pizza.Title)
	db.Create(&pizza)
	return c.JSON(pizza)
}

func DeletePizza(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var pizza Pizza
	db.First(&pizza, id)
	fmt.Println(pizza.Title)
	if pizza.Title == "" {
		return c.Status(500).SendString("No Pizza Found with ID")
	}
	db.Delete(&pizza)
	return c.SendString("Pizza Successfully deleted")
}
