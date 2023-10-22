package db

import "github.com/obynonwane/hotel-reservation/types"

type UserStore interface {
	GetUserById(string) (*types.User, error)
}

type MongoUserStore struct{}
