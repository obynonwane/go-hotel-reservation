package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/obynonwane/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Obinna",
		LastName:  "Johnson",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
