// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import types "github.com/SmartBFT-Go/consensus/v2/pkg/types"

// RequestsTimer is an autogenerated mock type for the RequestsTimer type
type RequestsTimer struct {
	mock.Mock
}

// RemoveRequest provides a mock function with given fields: request
func (_m *RequestsTimer) RemoveRequest(request types.RequestInfo) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.RequestInfo) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RestartTimers provides a mock function with given fields:
func (_m *RequestsTimer) RestartTimers() {
	_m.Called()
}

// StopTimers provides a mock function with given fields:
func (_m *RequestsTimer) StopTimers() {
	_m.Called()
}