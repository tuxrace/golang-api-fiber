package main

import (
	"fmt"
	"github.com/tuxrace/golang-api-fiber/appliances"
	"github.com/tuxrace/golang-api-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func hello(c *fiber.Ctx){
	c.Send("Hello")
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
	app.Get("/", hello)

	app.Get("/api/appliances", appliances.GetAppliances)
	app.Get("/api/appliances/:id", appliances.GetAppliance)
	app.Post("/api/appliances", appliances.NewAppliance)
	app.Delete("/api/appliances/:id", appliances.DeleteAppliance)
}

func main(){
	app := fiber.New()

	startDB()
	setupRoutes(app)
	app.Listen(3001)
	defer database.DB.Close()
}