package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAddress = "addresses"
)

type Address struct {
	ID     primitive.ObjectID `bson:"_id" json:"-"`
	Street string             `bson:"street" form:"street" binding:"required" json:"street"`
	City   string             `bson:"city" form:"city" binding:"required" json:"city"`
	UserID primitive.ObjectID `bson:"userID" json:"-"`
}

type AddressRepository interface {
	Create(c context.Context, address *Address) error
	FetchAddressByUserID(c context.Context, userID string) ([]Address, error)
}

type AddressUsecase interface {
	Create(c context.Context, address *Address) error
	FetchAddressByUserID(c context.Context, userID string) ([]Address, error)
}
