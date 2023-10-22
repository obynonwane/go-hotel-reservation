package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/obynonwane/hotel-reservation/db"
	"github.com/obynonwane/hotel-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

// create a constructor
func NewUserHandler() *UserHandler {

}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserById(id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Obinna",
		LastName:  "Johnson",
	}
	return c.JSON(u)
}
