// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/erikrios/reog-apps-apis/entity"
	mock "github.com/stretchr/testify/mock"
)

// AddressRepository is an autogenerated mock type for the AddressRepository type
type AddressRepository struct {
	mock.Mock
}

// Update provides a mock function with given fields: ctx, id, _a2
func (_m *AddressRepository) Update(ctx context.Context, id string, _a2 entity.Address) error {
	ret := _m.Called(ctx, id, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.Address) error); ok {
		r0 = rf(ctx, id, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
