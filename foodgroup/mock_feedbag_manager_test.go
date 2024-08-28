// Code generated by mockery v2.45.0. DO NOT EDIT.

package foodgroup

import (
	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	time "time"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockFeedbagManager is an autogenerated mock type for the FeedbagManager type
type mockFeedbagManager struct {
	mock.Mock
}

type mockFeedbagManager_Expecter struct {
	mock *mock.Mock
}

func (_m *mockFeedbagManager) EXPECT() *mockFeedbagManager_Expecter {
	return &mockFeedbagManager_Expecter{mock: &_m.Mock}
}

// AdjacentUsers provides a mock function with given fields: screenName
func (_m *mockFeedbagManager) AdjacentUsers(screenName state.IdentScreenName) ([]state.IdentScreenName, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for AdjacentUsers")
	}

	var r0 []state.IdentScreenName
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) ([]state.IdentScreenName, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) []state.IdentScreenName); ok {
		r0 = rf(screenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.IdentScreenName)
		}
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockFeedbagManager_AdjacentUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AdjacentUsers'
type mockFeedbagManager_AdjacentUsers_Call struct {
	*mock.Call
}

// AdjacentUsers is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockFeedbagManager_Expecter) AdjacentUsers(screenName interface{}) *mockFeedbagManager_AdjacentUsers_Call {
	return &mockFeedbagManager_AdjacentUsers_Call{Call: _e.mock.On("AdjacentUsers", screenName)}
}

func (_c *mockFeedbagManager_AdjacentUsers_Call) Run(run func(screenName state.IdentScreenName)) *mockFeedbagManager_AdjacentUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockFeedbagManager_AdjacentUsers_Call) Return(_a0 []state.IdentScreenName, _a1 error) *mockFeedbagManager_AdjacentUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockFeedbagManager_AdjacentUsers_Call) RunAndReturn(run func(state.IdentScreenName) ([]state.IdentScreenName, error)) *mockFeedbagManager_AdjacentUsers_Call {
	_c.Call.Return(run)
	return _c
}

// BlockedState provides a mock function with given fields: screenName1, screenName2
func (_m *mockFeedbagManager) BlockedState(screenName1 state.IdentScreenName, screenName2 state.IdentScreenName) (state.BlockedState, error) {
	ret := _m.Called(screenName1, screenName2)

	if len(ret) == 0 {
		panic("no return value specified for BlockedState")
	}

	var r0 state.BlockedState
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, state.IdentScreenName) (state.BlockedState, error)); ok {
		return rf(screenName1, screenName2)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, state.IdentScreenName) state.BlockedState); ok {
		r0 = rf(screenName1, screenName2)
	} else {
		r0 = ret.Get(0).(state.BlockedState)
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName, state.IdentScreenName) error); ok {
		r1 = rf(screenName1, screenName2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockFeedbagManager_BlockedState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockedState'
type mockFeedbagManager_BlockedState_Call struct {
	*mock.Call
}

// BlockedState is a helper method to define mock.On call
//   - screenName1 state.IdentScreenName
//   - screenName2 state.IdentScreenName
func (_e *mockFeedbagManager_Expecter) BlockedState(screenName1 interface{}, screenName2 interface{}) *mockFeedbagManager_BlockedState_Call {
	return &mockFeedbagManager_BlockedState_Call{Call: _e.mock.On("BlockedState", screenName1, screenName2)}
}

func (_c *mockFeedbagManager_BlockedState_Call) Run(run func(screenName1 state.IdentScreenName, screenName2 state.IdentScreenName)) *mockFeedbagManager_BlockedState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockFeedbagManager_BlockedState_Call) Return(_a0 state.BlockedState, _a1 error) *mockFeedbagManager_BlockedState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockFeedbagManager_BlockedState_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName) (state.BlockedState, error)) *mockFeedbagManager_BlockedState_Call {
	_c.Call.Return(run)
	return _c
}

// Buddies provides a mock function with given fields: screenName
func (_m *mockFeedbagManager) Buddies(screenName state.IdentScreenName) ([]state.IdentScreenName, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for Buddies")
	}

	var r0 []state.IdentScreenName
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) ([]state.IdentScreenName, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) []state.IdentScreenName); ok {
		r0 = rf(screenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.IdentScreenName)
		}
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockFeedbagManager_Buddies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Buddies'
type mockFeedbagManager_Buddies_Call struct {
	*mock.Call
}

// Buddies is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockFeedbagManager_Expecter) Buddies(screenName interface{}) *mockFeedbagManager_Buddies_Call {
	return &mockFeedbagManager_Buddies_Call{Call: _e.mock.On("Buddies", screenName)}
}

func (_c *mockFeedbagManager_Buddies_Call) Run(run func(screenName state.IdentScreenName)) *mockFeedbagManager_Buddies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockFeedbagManager_Buddies_Call) Return(_a0 []state.IdentScreenName, _a1 error) *mockFeedbagManager_Buddies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockFeedbagManager_Buddies_Call) RunAndReturn(run func(state.IdentScreenName) ([]state.IdentScreenName, error)) *mockFeedbagManager_Buddies_Call {
	_c.Call.Return(run)
	return _c
}

// Feedbag provides a mock function with given fields: screenName
func (_m *mockFeedbagManager) Feedbag(screenName state.IdentScreenName) ([]wire.FeedbagItem, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for Feedbag")
	}

	var r0 []wire.FeedbagItem
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) ([]wire.FeedbagItem, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) []wire.FeedbagItem); ok {
		r0 = rf(screenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]wire.FeedbagItem)
		}
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockFeedbagManager_Feedbag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Feedbag'
type mockFeedbagManager_Feedbag_Call struct {
	*mock.Call
}

// Feedbag is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockFeedbagManager_Expecter) Feedbag(screenName interface{}) *mockFeedbagManager_Feedbag_Call {
	return &mockFeedbagManager_Feedbag_Call{Call: _e.mock.On("Feedbag", screenName)}
}

func (_c *mockFeedbagManager_Feedbag_Call) Run(run func(screenName state.IdentScreenName)) *mockFeedbagManager_Feedbag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockFeedbagManager_Feedbag_Call) Return(_a0 []wire.FeedbagItem, _a1 error) *mockFeedbagManager_Feedbag_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockFeedbagManager_Feedbag_Call) RunAndReturn(run func(state.IdentScreenName) ([]wire.FeedbagItem, error)) *mockFeedbagManager_Feedbag_Call {
	_c.Call.Return(run)
	return _c
}

// FeedbagDelete provides a mock function with given fields: screenName, items
func (_m *mockFeedbagManager) FeedbagDelete(screenName state.IdentScreenName, items []wire.FeedbagItem) error {
	ret := _m.Called(screenName, items)

	if len(ret) == 0 {
		panic("no return value specified for FeedbagDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, []wire.FeedbagItem) error); ok {
		r0 = rf(screenName, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockFeedbagManager_FeedbagDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeedbagDelete'
type mockFeedbagManager_FeedbagDelete_Call struct {
	*mock.Call
}

// FeedbagDelete is a helper method to define mock.On call
//   - screenName state.IdentScreenName
//   - items []wire.FeedbagItem
func (_e *mockFeedbagManager_Expecter) FeedbagDelete(screenName interface{}, items interface{}) *mockFeedbagManager_FeedbagDelete_Call {
	return &mockFeedbagManager_FeedbagDelete_Call{Call: _e.mock.On("FeedbagDelete", screenName, items)}
}

func (_c *mockFeedbagManager_FeedbagDelete_Call) Run(run func(screenName state.IdentScreenName, items []wire.FeedbagItem)) *mockFeedbagManager_FeedbagDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].([]wire.FeedbagItem))
	})
	return _c
}

func (_c *mockFeedbagManager_FeedbagDelete_Call) Return(_a0 error) *mockFeedbagManager_FeedbagDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFeedbagManager_FeedbagDelete_Call) RunAndReturn(run func(state.IdentScreenName, []wire.FeedbagItem) error) *mockFeedbagManager_FeedbagDelete_Call {
	_c.Call.Return(run)
	return _c
}

// FeedbagLastModified provides a mock function with given fields: screenName
func (_m *mockFeedbagManager) FeedbagLastModified(screenName state.IdentScreenName) (time.Time, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for FeedbagLastModified")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) (time.Time, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) time.Time); ok {
		r0 = rf(screenName)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockFeedbagManager_FeedbagLastModified_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeedbagLastModified'
type mockFeedbagManager_FeedbagLastModified_Call struct {
	*mock.Call
}

// FeedbagLastModified is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockFeedbagManager_Expecter) FeedbagLastModified(screenName interface{}) *mockFeedbagManager_FeedbagLastModified_Call {
	return &mockFeedbagManager_FeedbagLastModified_Call{Call: _e.mock.On("FeedbagLastModified", screenName)}
}

func (_c *mockFeedbagManager_FeedbagLastModified_Call) Run(run func(screenName state.IdentScreenName)) *mockFeedbagManager_FeedbagLastModified_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockFeedbagManager_FeedbagLastModified_Call) Return(_a0 time.Time, _a1 error) *mockFeedbagManager_FeedbagLastModified_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockFeedbagManager_FeedbagLastModified_Call) RunAndReturn(run func(state.IdentScreenName) (time.Time, error)) *mockFeedbagManager_FeedbagLastModified_Call {
	_c.Call.Return(run)
	return _c
}

// FeedbagUpsert provides a mock function with given fields: screenName, items
func (_m *mockFeedbagManager) FeedbagUpsert(screenName state.IdentScreenName, items []wire.FeedbagItem) error {
	ret := _m.Called(screenName, items)

	if len(ret) == 0 {
		panic("no return value specified for FeedbagUpsert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, []wire.FeedbagItem) error); ok {
		r0 = rf(screenName, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockFeedbagManager_FeedbagUpsert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FeedbagUpsert'
type mockFeedbagManager_FeedbagUpsert_Call struct {
	*mock.Call
}

// FeedbagUpsert is a helper method to define mock.On call
//   - screenName state.IdentScreenName
//   - items []wire.FeedbagItem
func (_e *mockFeedbagManager_Expecter) FeedbagUpsert(screenName interface{}, items interface{}) *mockFeedbagManager_FeedbagUpsert_Call {
	return &mockFeedbagManager_FeedbagUpsert_Call{Call: _e.mock.On("FeedbagUpsert", screenName, items)}
}

func (_c *mockFeedbagManager_FeedbagUpsert_Call) Run(run func(screenName state.IdentScreenName, items []wire.FeedbagItem)) *mockFeedbagManager_FeedbagUpsert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].([]wire.FeedbagItem))
	})
	return _c
}

func (_c *mockFeedbagManager_FeedbagUpsert_Call) Return(_a0 error) *mockFeedbagManager_FeedbagUpsert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFeedbagManager_FeedbagUpsert_Call) RunAndReturn(run func(state.IdentScreenName, []wire.FeedbagItem) error) *mockFeedbagManager_FeedbagUpsert_Call {
	_c.Call.Return(run)
	return _c
}

// newMockFeedbagManager creates a new instance of mockFeedbagManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockFeedbagManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockFeedbagManager {
	mock := &mockFeedbagManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
