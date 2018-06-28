// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import slack "github.com/nlopes/slack"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetUserByEmail provides a mock function with given fields: _a0
func (_m *Client) GetUserByEmail(_a0 string) (*slack.User, error) {
	ret := _m.Called(_a0)

	var r0 *slack.User
	if rf, ok := ret.Get(0).(func(string) *slack.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slack.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JoinChannelContext provides a mock function with given fields: _a0, _a1
func (_m *Client) JoinChannelContext(_a0 context.Context, _a1 string) (*slack.Channel, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *slack.Channel
	if rf, ok := ret.Get(0).(func(context.Context, string) *slack.Channel); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slack.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenIMChannel provides a mock function with given fields: _a0
func (_m *Client) OpenIMChannel(_a0 string) (bool, bool, string, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(string) string); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Get(2).(string)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string) error); ok {
		r3 = rf(_a0)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// SetDebug provides a mock function with given fields: _a0
func (_m *Client) SetDebug(_a0 bool) {
	_m.Called(_a0)
}