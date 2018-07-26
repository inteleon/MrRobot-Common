// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import slack "github.com/nlopes/slack"

// RTM is an autogenerated mock type for the RTM type
type RTM struct {
	mock.Mock
}

// GetIncomingEvents provides a mock function with given fields:
func (_m *RTM) GetIncomingEvents() chan slack.RTMEvent {
	ret := _m.Called()

	var r0 chan slack.RTMEvent
	if rf, ok := ret.Get(0).(func() chan slack.RTMEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan slack.RTMEvent)
		}
	}

	return r0
}

// GetInfo provides a mock function with given fields:
func (_m *RTM) GetInfo() *slack.Info {
	ret := _m.Called()

	var r0 *slack.Info
	if rf, ok := ret.Get(0).(func() *slack.Info); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slack.Info)
		}
	}

	return r0
}

// ManageConnection provides a mock function with given fields:
func (_m *RTM) ManageConnection() {
	_m.Called()
}

// NewOutgoingMessage provides a mock function with given fields: _a0, _a1, _a2
func (_m *RTM) NewOutgoingMessage(_a0 string, _a1 string, _a2 ...slack.RTMsgOption) *slack.OutgoingMessage {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *slack.OutgoingMessage
	if rf, ok := ret.Get(0).(func(string, string, ...slack.RTMsgOption) *slack.OutgoingMessage); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slack.OutgoingMessage)
		}
	}

	return r0
}

// SendMessage provides a mock function with given fields: _a0
func (_m *RTM) SendMessage(_a0 *slack.OutgoingMessage) {
	_m.Called(_a0)
}
