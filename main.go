package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/obynonwane/hotel-reservation/api"
	"github.com/obynonwane/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)

	users := types.User{
		FirstName: "Obinna",
		LastName:  "Johnson",
	}

	_, err = coll.InsertOne(ctx, users)
	if err != nil {
		log.Fatal(err)
	}

	var obinna types.User
	if err := coll.FindOne(ctx, bson.M{}).Decode(&obinna); err != nil {
		log.Fatal(err)
	}

	fmt.Println(obinna)

	listenAddress := flag.String("listenAddr", ":3000", "The listen address of the api server")
	flag.Parse()
	app := fiber.New()

	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddress)
}
