package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/obynonwane/hotel-reservation/db"
)

type UserHandler struct {
	userStore db.UserStore
}

// create a constructor function
func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(nil)
}
