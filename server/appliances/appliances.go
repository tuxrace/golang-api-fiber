package appliances

import (
	// "fmt"

	"github.com/tuxrace/golang-api-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Appliances struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetAppliances(c *fiber.Ctx) {
	db := database.DB
	var appliances []Appliances
	db.Find(&appliances)
	c.JSON(appliances)
}

func GetAppliance(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var appliances Appliances
	db.Find(&appliances, id)
	c.JSON(appliances)
}

func NewAppliance(c *fiber.Ctx) {
	db := database.DB
	appliances := new(Appliances)
	if err := c.BodyParser(appliances); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&appliances)
}

func DeleteAppliance(c *fiber.Ctx) {
	c.Send("Delete Appliances")
}
