package main

import (
	"github.com/gofiber/fiber",
	"github.com/tuxrace/golang-api-fiber/appliances"
)

func hello(c *fiber.Ctx){
	c.Send("Hello")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/appliances", appliances.GetBooks)
	app.Get("/api/v1/appliances/:id", appliances.GetBook)
	app.Post("/api/v1/appliances", appliances.NewBook)
	app.Delete("/api/v1/appliances/:id", appliances.DeleteBook)
}

func startDB() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "appliances.db")
	if err != nil {
		panic("Error connect")
	}
	fmt.Println("Connected")
	database.DBConn.AutoMigrate();
}

func main(){
	app := fiber.New()
	app.Get("/", hello)
	app.Listen(3001)
}