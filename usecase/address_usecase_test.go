package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain/mocks"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func TestFetchAddressByUserID(t *testing.T) {
	mockAddressRepository := new(mocks.AddressRepository)
	userObjectID := primitive.NewObjectID()
	userID := userObjectID.Hex()

	t.Run("success", func(t *testing.T) {

		mockAddress := domain.Address{
			ID:     primitive.NewObjectID(),
			Street: "Test Street",
			City: "Test City",
			UserID: userObjectID,
		}


		mockListAddress := make([]domain.Address, 0)
		mockListAddress = append(mockListAddress, mockAddress)

		mockAddressRepository.On("FetchAddressByUserID", mock.Anything, userID).Return(mockListAddress, nil).Once()

		u := usecase.NewAddressUsecase(mockAddressRepository, time.Second*2)

		list, err := u.FetchAddressByUserID(context.Background(), userID)


		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockListAddress))

		mockAddressRepository.AssertExpectations(t)
	})


	t.Run("error", func(t *testing.T) {
		mockAddressRepository.On("FetchAddressByUserID", mock.Anything, userID).Return(nil, errors.New("Unexpected")).Once()

		u := usecase.NewAddressUsecase(mockAddressRepository, time.Second*2)

		list, err := u.FetchAddressByUserID(context.Background(), userID)

		assert.Error(t, err)
		assert.Nil(t, list)

		mockAddressRepository.AssertExpectations(t)
	})
}
