package main

import (
	"fmt"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/tuxrace/golang-api-fiber/appliances"
	"github.com/tuxrace/golang-api-fiber/database"
)

func index(c *fiber.Ctx) {
	c.Send("Index")
}

func startDB() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "appliances.db")

	if err != nil {
		panic("Error connect")
	}

	fmt.Println("Connected")
	database.DB.AutoMigrate(&appliances.Appliances{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/", index)

	app.Get("/api/appliances", appliances.GetAppliances)
	app.Get("/api/appliances/:id", appliances.GetAppliance)
	app.Get("/api/appliances-search", appliances.GetAppliancesSearch)
	app.Post("/api/appliances", appliances.NewAppliance)
	app.Delete("/api/appliances/:id", appliances.DeleteAppliance)
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	startDB()
	setupRoutes(app)
	app.Listen(3001)
	defer database.DB.Close()
}
