package appliances

import (
	"fmt"

	"github.com/tuxrace/golang-api-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Appliances struct {
	gorm.Model
	SerialNumber  string `json:"serialNumber"`
	Brand string `json:"brand"`
	ModelName string `json:"model"`
	Status string    `json:"status"`
	DateBought string    `json:"dateBought"`

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

func GetAppliancesSearch(c *fiber.Ctx) {
	category := c.Query("category")
	search := c.Query("search")
	categoryParam := fmt.Sprintf("%s = ?", category)
	db := database.DB
	var appliances Appliances
	// db.Find(&appliances, id)
	db.Where(categoryParam, search).Find(&appliances)
	c.JSON(appliances)
}

func NewAppliance(c *fiber.Ctx) {
	db := database.DB
	// appliances := new(Appliances)
	// if err := c.BodyParser(appliances); err != nil {
	// 	c.Status(503).Send(err)
	// 	return
	// }
	var appliances Appliances
	appliances.SerialNumber = "1234"
	appliances.Brand = "Panasonic"
	appliances.ModelName = "Tub"
	appliances.Status = "Sold"
	appliances.DateBought = "June 20, 2020"
	db.Create(&appliances)
	c.JSON(appliances)
}

func DeleteAppliance(c *fiber.Ctx) {
	c.Send("Delete Appliances")
}
