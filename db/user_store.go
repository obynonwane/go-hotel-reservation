package db

import (
	"context"

	"github.com/obynonwane/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

// Specifies methods that any type can implement
type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

// this struct is intended to implement UserStore interface
type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

/*
	constructor function - creates an instance of a MongoUserStore struct

When called, this function initializes a MongoUserStore
with the provided MongoDB client and returns a pointer
to the initialized struct. This allows other parts of
the code to create instances of MongoUserStore for
working with MongoDB user data.
*
*/
func NewMongoUserStore(client *mongo.Client) *MongoUserStore {

	//returns a pointer to MongoUserStore
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(userColl),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
