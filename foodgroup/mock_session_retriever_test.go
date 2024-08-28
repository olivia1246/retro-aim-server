// Code generated by mockery v2.45.0. DO NOT EDIT.

package foodgroup

import (
	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"
)

// mockSessionRetriever is an autogenerated mock type for the SessionRetriever type
type mockSessionRetriever struct {
	mock.Mock
}

type mockSessionRetriever_Expecter struct {
	mock *mock.Mock
}

func (_m *mockSessionRetriever) EXPECT() *mockSessionRetriever_Expecter {
	return &mockSessionRetriever_Expecter{mock: &_m.Mock}
}

// RetrieveSession provides a mock function with given fields: screenName
func (_m *mockSessionRetriever) RetrieveSession(screenName state.IdentScreenName) *state.Session {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveSession")
	}

	var r0 *state.Session
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) *state.Session); ok {
		r0 = rf(screenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Session)
		}
	}

	return r0
}

// mockSessionRetriever_RetrieveSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveSession'
type mockSessionRetriever_RetrieveSession_Call struct {
	*mock.Call
}

// RetrieveSession is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockSessionRetriever_Expecter) RetrieveSession(screenName interface{}) *mockSessionRetriever_RetrieveSession_Call {
	return &mockSessionRetriever_RetrieveSession_Call{Call: _e.mock.On("RetrieveSession", screenName)}
}

func (_c *mockSessionRetriever_RetrieveSession_Call) Run(run func(screenName state.IdentScreenName)) *mockSessionRetriever_RetrieveSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockSessionRetriever_RetrieveSession_Call) Return(_a0 *state.Session) *mockSessionRetriever_RetrieveSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockSessionRetriever_RetrieveSession_Call) RunAndReturn(run func(state.IdentScreenName) *state.Session) *mockSessionRetriever_RetrieveSession_Call {
	_c.Call.Return(run)
	return _c
}

// newMockSessionRetriever creates a new instance of mockSessionRetriever. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSessionRetriever(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockSessionRetriever {
	mock := &mockSessionRetriever{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
