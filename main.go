package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/obynonwane/hotel-reservation/api"
)

func main() {
	listenAddress := flag.String("listenAddr", ":3000", "The listen address of the api server")
	flag.Parse()
	app := fiber.New()

	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddress)
}
