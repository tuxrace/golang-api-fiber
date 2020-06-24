package appliances

import (
	"github.com/gofiber/fiber"
)

func GetAppliances(c *fiber.Ctx) {
	c.Send("All Appliances")
}

func GetAppliances(c *fiber.Ctx) {
	c.Send("Single Appliances")
}

func NewAppliances(c *fiber.Ctx) {
	c.Send("New Appliances")
}

func DeleteAppliances(c *fiber.Ctx) {
	c.Send("Delete Appliances")
}
