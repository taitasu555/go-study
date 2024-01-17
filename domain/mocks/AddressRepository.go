// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	mock "github.com/stretchr/testify/mock"
)

// AddressRepository is an autogenerated mock type for the AddressRepository type
type AddressRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, address
func (_m *AddressRepository) Create(c context.Context, address *domain.Address) error {
	ret := _m.Called(c, address)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Address) error); ok {
		r0 = rf(c, address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchAddressByUserID provides a mock function with given fields: c, userID
func (_m *AddressRepository) FetchAddressByUserID(c context.Context, userID string) ([]domain.Address, error) {
	ret := _m.Called(c, userID)

	if len(ret) == 0 {
		panic("no return value specified for FetchAddressByUserID")
	}

	var r0 []domain.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.Address, error)); ok {
		return rf(c, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Address); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAddressRepository creates a new instance of AddressRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAddressRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AddressRepository {
	mock := &AddressRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}