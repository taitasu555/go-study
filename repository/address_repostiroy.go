package repository

import (
	"context"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressRepository struct {
	database   mongo.Database
	collection string
}

func NewAddressRepository(db mongo.Database, collection string) *AddressRepository {
	return &AddressRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *AddressRepository) Create(c context.Context, task *domain.Address) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, task)

	return err
}

func (tr *AddressRepository) FetchAddressByUserID(c context.Context, userID string) ([]domain.Address, error) {
	collection := tr.database.Collection(tr.collection)

	var address []domain.Address

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return address, err
	}
	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &address)
	if address == nil {
		return []domain.Address{}, err
	}

	return address, err
}
