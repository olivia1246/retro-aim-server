// Code generated by mockery v2.45.0. DO NOT EDIT.

package foodgroup

import (
	context "context"

	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockChatMessageRelayer is an autogenerated mock type for the ChatMessageRelayer type
type mockChatMessageRelayer struct {
	mock.Mock
}

type mockChatMessageRelayer_Expecter struct {
	mock *mock.Mock
}

func (_m *mockChatMessageRelayer) EXPECT() *mockChatMessageRelayer_Expecter {
	return &mockChatMessageRelayer_Expecter{mock: &_m.Mock}
}

// AllSessions provides a mock function with given fields: chatCookie
func (_m *mockChatMessageRelayer) AllSessions(chatCookie string) []*state.Session {
	ret := _m.Called(chatCookie)

	if len(ret) == 0 {
		panic("no return value specified for AllSessions")
	}

	var r0 []*state.Session
	if rf, ok := ret.Get(0).(func(string) []*state.Session); ok {
		r0 = rf(chatCookie)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.Session)
		}
	}

	return r0
}

// mockChatMessageRelayer_AllSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllSessions'
type mockChatMessageRelayer_AllSessions_Call struct {
	*mock.Call
}

// AllSessions is a helper method to define mock.On call
//   - chatCookie string
func (_e *mockChatMessageRelayer_Expecter) AllSessions(chatCookie interface{}) *mockChatMessageRelayer_AllSessions_Call {
	return &mockChatMessageRelayer_AllSessions_Call{Call: _e.mock.On("AllSessions", chatCookie)}
}

func (_c *mockChatMessageRelayer_AllSessions_Call) Run(run func(chatCookie string)) *mockChatMessageRelayer_AllSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *mockChatMessageRelayer_AllSessions_Call) Return(_a0 []*state.Session) *mockChatMessageRelayer_AllSessions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockChatMessageRelayer_AllSessions_Call) RunAndReturn(run func(string) []*state.Session) *mockChatMessageRelayer_AllSessions_Call {
	_c.Call.Return(run)
	return _c
}

// RelayToAllExcept provides a mock function with given fields: ctx, chatCookie, except, msg
func (_m *mockChatMessageRelayer) RelayToAllExcept(ctx context.Context, chatCookie string, except state.IdentScreenName, msg wire.SNACMessage) {
	_m.Called(ctx, chatCookie, except, msg)
}

// mockChatMessageRelayer_RelayToAllExcept_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RelayToAllExcept'
type mockChatMessageRelayer_RelayToAllExcept_Call struct {
	*mock.Call
}

// RelayToAllExcept is a helper method to define mock.On call
//   - ctx context.Context
//   - chatCookie string
//   - except state.IdentScreenName
//   - msg wire.SNACMessage
func (_e *mockChatMessageRelayer_Expecter) RelayToAllExcept(ctx interface{}, chatCookie interface{}, except interface{}, msg interface{}) *mockChatMessageRelayer_RelayToAllExcept_Call {
	return &mockChatMessageRelayer_RelayToAllExcept_Call{Call: _e.mock.On("RelayToAllExcept", ctx, chatCookie, except, msg)}
}

func (_c *mockChatMessageRelayer_RelayToAllExcept_Call) Run(run func(ctx context.Context, chatCookie string, except state.IdentScreenName, msg wire.SNACMessage)) *mockChatMessageRelayer_RelayToAllExcept_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(state.IdentScreenName), args[3].(wire.SNACMessage))
	})
	return _c
}

func (_c *mockChatMessageRelayer_RelayToAllExcept_Call) Return() *mockChatMessageRelayer_RelayToAllExcept_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockChatMessageRelayer_RelayToAllExcept_Call) RunAndReturn(run func(context.Context, string, state.IdentScreenName, wire.SNACMessage)) *mockChatMessageRelayer_RelayToAllExcept_Call {
	_c.Call.Return(run)
	return _c
}

// RelayToScreenName provides a mock function with given fields: ctx, chatCookie, recipient, msg
func (_m *mockChatMessageRelayer) RelayToScreenName(ctx context.Context, chatCookie string, recipient state.IdentScreenName, msg wire.SNACMessage) {
	_m.Called(ctx, chatCookie, recipient, msg)
}

// mockChatMessageRelayer_RelayToScreenName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RelayToScreenName'
type mockChatMessageRelayer_RelayToScreenName_Call struct {
	*mock.Call
}

// RelayToScreenName is a helper method to define mock.On call
//   - ctx context.Context
//   - chatCookie string
//   - recipient state.IdentScreenName
//   - msg wire.SNACMessage
func (_e *mockChatMessageRelayer_Expecter) RelayToScreenName(ctx interface{}, chatCookie interface{}, recipient interface{}, msg interface{}) *mockChatMessageRelayer_RelayToScreenName_Call {
	return &mockChatMessageRelayer_RelayToScreenName_Call{Call: _e.mock.On("RelayToScreenName", ctx, chatCookie, recipient, msg)}
}

func (_c *mockChatMessageRelayer_RelayToScreenName_Call) Run(run func(ctx context.Context, chatCookie string, recipient state.IdentScreenName, msg wire.SNACMessage)) *mockChatMessageRelayer_RelayToScreenName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(state.IdentScreenName), args[3].(wire.SNACMessage))
	})
	return _c
}

func (_c *mockChatMessageRelayer_RelayToScreenName_Call) Return() *mockChatMessageRelayer_RelayToScreenName_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockChatMessageRelayer_RelayToScreenName_Call) RunAndReturn(run func(context.Context, string, state.IdentScreenName, wire.SNACMessage)) *mockChatMessageRelayer_RelayToScreenName_Call {
	_c.Call.Return(run)
	return _c
}

// newMockChatMessageRelayer creates a new instance of mockChatMessageRelayer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockChatMessageRelayer(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockChatMessageRelayer {
	mock := &mockChatMessageRelayer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
