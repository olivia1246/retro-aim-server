// Code generated by mockery v2.45.0. DO NOT EDIT.

package http

import (
	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"
)

// mockUserManager is an autogenerated mock type for the UserManager type
type mockUserManager struct {
	mock.Mock
}

type mockUserManager_Expecter struct {
	mock *mock.Mock
}

func (_m *mockUserManager) EXPECT() *mockUserManager_Expecter {
	return &mockUserManager_Expecter{mock: &_m.Mock}
}

// AllUsers provides a mock function with given fields:
func (_m *mockUserManager) AllUsers() ([]state.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AllUsers")
	}

	var r0 []state.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]state.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []state.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserManager_AllUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllUsers'
type mockUserManager_AllUsers_Call struct {
	*mock.Call
}

// AllUsers is a helper method to define mock.On call
func (_e *mockUserManager_Expecter) AllUsers() *mockUserManager_AllUsers_Call {
	return &mockUserManager_AllUsers_Call{Call: _e.mock.On("AllUsers")}
}

func (_c *mockUserManager_AllUsers_Call) Run(run func()) *mockUserManager_AllUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockUserManager_AllUsers_Call) Return(_a0 []state.User, _a1 error) *mockUserManager_AllUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserManager_AllUsers_Call) RunAndReturn(run func() ([]state.User, error)) *mockUserManager_AllUsers_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: screenName
func (_m *mockUserManager) DeleteUser(screenName state.IdentScreenName) error {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) error); ok {
		r0 = rf(screenName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockUserManager_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type mockUserManager_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockUserManager_Expecter) DeleteUser(screenName interface{}) *mockUserManager_DeleteUser_Call {
	return &mockUserManager_DeleteUser_Call{Call: _e.mock.On("DeleteUser", screenName)}
}

func (_c *mockUserManager_DeleteUser_Call) Run(run func(screenName state.IdentScreenName)) *mockUserManager_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockUserManager_DeleteUser_Call) Return(_a0 error) *mockUserManager_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockUserManager_DeleteUser_Call) RunAndReturn(run func(state.IdentScreenName) error) *mockUserManager_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// InsertUser provides a mock function with given fields: u
func (_m *mockUserManager) InsertUser(u state.User) error {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for InsertUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.User) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockUserManager_InsertUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertUser'
type mockUserManager_InsertUser_Call struct {
	*mock.Call
}

// InsertUser is a helper method to define mock.On call
//   - u state.User
func (_e *mockUserManager_Expecter) InsertUser(u interface{}) *mockUserManager_InsertUser_Call {
	return &mockUserManager_InsertUser_Call{Call: _e.mock.On("InsertUser", u)}
}

func (_c *mockUserManager_InsertUser_Call) Run(run func(u state.User)) *mockUserManager_InsertUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.User))
	})
	return _c
}

func (_c *mockUserManager_InsertUser_Call) Return(_a0 error) *mockUserManager_InsertUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockUserManager_InsertUser_Call) RunAndReturn(run func(state.User) error) *mockUserManager_InsertUser_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserPassword provides a mock function with given fields: screenName, newPassword
func (_m *mockUserManager) SetUserPassword(screenName state.IdentScreenName, newPassword string) error {
	ret := _m.Called(screenName, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for SetUserPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, string) error); ok {
		r0 = rf(screenName, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockUserManager_SetUserPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserPassword'
type mockUserManager_SetUserPassword_Call struct {
	*mock.Call
}

// SetUserPassword is a helper method to define mock.On call
//   - screenName state.IdentScreenName
//   - newPassword string
func (_e *mockUserManager_Expecter) SetUserPassword(screenName interface{}, newPassword interface{}) *mockUserManager_SetUserPassword_Call {
	return &mockUserManager_SetUserPassword_Call{Call: _e.mock.On("SetUserPassword", screenName, newPassword)}
}

func (_c *mockUserManager_SetUserPassword_Call) Run(run func(screenName state.IdentScreenName, newPassword string)) *mockUserManager_SetUserPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(string))
	})
	return _c
}

func (_c *mockUserManager_SetUserPassword_Call) Return(_a0 error) *mockUserManager_SetUserPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockUserManager_SetUserPassword_Call) RunAndReturn(run func(state.IdentScreenName, string) error) *mockUserManager_SetUserPassword_Call {
	_c.Call.Return(run)
	return _c
}

// User provides a mock function with given fields: screenName
func (_m *mockUserManager) User(screenName state.IdentScreenName) (*state.User, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for User")
	}

	var r0 *state.User
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) (*state.User, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) *state.User); ok {
		r0 = rf(screenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.User)
		}
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserManager_User_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'User'
type mockUserManager_User_Call struct {
	*mock.Call
}

// User is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockUserManager_Expecter) User(screenName interface{}) *mockUserManager_User_Call {
	return &mockUserManager_User_Call{Call: _e.mock.On("User", screenName)}
}

func (_c *mockUserManager_User_Call) Run(run func(screenName state.IdentScreenName)) *mockUserManager_User_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockUserManager_User_Call) Return(_a0 *state.User, _a1 error) *mockUserManager_User_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserManager_User_Call) RunAndReturn(run func(state.IdentScreenName) (*state.User, error)) *mockUserManager_User_Call {
	_c.Call.Return(run)
	return _c
}

// newMockUserManager creates a new instance of mockUserManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockUserManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockUserManager {
	mock := &mockUserManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
