// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	payload "github.com/erikrios/reog-apps-apis/model/payload"

	response "github.com/erikrios/reog-apps-apis/model/response"
)

// GroupService is an autogenerated mock type for the GroupService type
type GroupService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, p
func (_m *GroupService) Create(ctx context.Context, p payload.CreateGroup) (string, error) {
	ret := _m.Called(ctx, p)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, payload.CreateGroup) string); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, payload.CreateGroup) error); ok {
		r1 = rf(ctx, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *GroupService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateQRCode provides a mock function with given fields: ctx, id
func (_m *GroupService) GenerateQRCode(ctx context.Context, id string) ([]byte, error) {
	ret := _m.Called(ctx, id)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *GroupService) GetAll(ctx context.Context) ([]response.Group, error) {
	ret := _m.Called(ctx)

	var r0 []response.Group
	if rf, ok := ret.Get(0).(func(context.Context) []response.Group); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *GroupService) GetByID(ctx context.Context, id string) (response.Group, error) {
	ret := _m.Called(ctx, id)

	var r0 response.Group
	if rf, ok := ret.Get(0).(func(context.Context, string) response.Group); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(response.Group)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, p
func (_m *GroupService) Update(ctx context.Context, id string, p payload.UpdateGroup) error {
	ret := _m.Called(ctx, id, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, payload.UpdateGroup) error); ok {
		r0 = rf(ctx, id, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
