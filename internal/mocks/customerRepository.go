// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/hinccvi/saga/internal/entity"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CustomerRepository is an autogenerated mock type for the Repository type
type CustomerRepository struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: ctx, c
func (_m *CustomerRepository) CreateCustomer(ctx context.Context, c entity.Customer) (uuid.UUID, error) {
	ret := _m.Called(ctx, c)

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func(context.Context, entity.Customer) uuid.UUID); ok {
		r0 = rf(ctx, c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.Customer) error); ok {
		r1 = rf(ctx, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCustomerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCustomerRepository creates a new instance of CustomerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCustomerRepository(t mockConstructorTestingTNewCustomerRepository) *CustomerRepository {
	mock := &CustomerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
