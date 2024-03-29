// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	qrcode "github.com/skip2/go-qrcode"
	mock "github.com/stretchr/testify/mock"
)

// QRCodeGenerator is an autogenerated mock type for the QRCodeGenerator type
type QRCodeGenerator struct {
	mock.Mock
}

// GenerateQRCode provides a mock function with given fields: content, level, size
func (_m *QRCodeGenerator) GenerateQRCode(content string, level qrcode.RecoveryLevel, size int) ([]byte, error) {
	ret := _m.Called(content, level, size)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, qrcode.RecoveryLevel, int) []byte); ok {
		r0 = rf(content, level, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, qrcode.RecoveryLevel, int) error); ok {
		r1 = rf(content, level, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
