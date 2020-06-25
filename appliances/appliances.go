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
	SerialNumber string `json:"serialNumber"`
	Brand string `json:"brand"`
	ModelName string `json:"model"`
	Status string `json:"status"`
	DateBought string `json:"dateBought"`
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
	categoryParam := fmt.Sprintf("%s LIKE ?", category)
	searchParam := search+"%"
	db := database.DB
	var appliances []Appliances
	db.Where(categoryParam, searchParam).Find(&appliances)
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
	appliances.SerialNumber = "9101"
	appliances.Brand = "Philips"
	appliances.ModelName = "Wall Lamp"
	appliances.Status = "Sold"
	appliances.DateBought = "June 26, 2020"
	db.Create(&appliances)
	c.JSON(appliances)
}

func DeleteAppliance(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	var appliances Appliances
	db.First(&appliances, id)
	if appliances.SerialNumber == "" {
        c.Status(500).Send("No appliances Found with SerialNumber")
        return
	}
	db.Delete(&appliances)
	c.Send("appliances Successfully deleted")
}
