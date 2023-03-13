// Code generated by mockery v2.20.2. DO NOT EDIT.

package privatemessagingmocks

import (
	context "context"

	ffapi "github.com/hyperledger/firefly-common/pkg/ffapi"
	core "github.com/hyperledger/firefly/pkg/core"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"

	syncasync "github.com/hyperledger/firefly/internal/syncasync"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// EnsureLocalGroup provides a mock function with given fields: ctx, group, creator
func (_m *Manager) EnsureLocalGroup(ctx context.Context, group *core.Group, creator *core.Member) (bool, error) {
	ret := _m.Called(ctx, group, creator)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Group, *core.Member) (bool, error)); ok {
		return rf(ctx, group, creator)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Group, *core.Member) bool); ok {
		r0 = rf(ctx, group, creator)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Group, *core.Member) error); ok {
		r1 = rf(ctx, group, creator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupByID provides a mock function with given fields: ctx, id
func (_m *Manager) GetGroupByID(ctx context.Context, id string) (*core.Group, error) {
	ret := _m.Called(ctx, id)

	var r0 *core.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*core.Group, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Group); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroups provides a mock function with given fields: ctx, filter
func (_m *Manager) GetGroups(ctx context.Context, filter ffapi.AndFilter) ([]*core.Group, *ffapi.FilterResult, error) {
	ret := _m.Called(ctx, filter)

	var r0 []*core.Group
	var r1 *ffapi.FilterResult
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) ([]*core.Group, *ffapi.FilterResult, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ffapi.AndFilter) []*core.Group); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ffapi.AndFilter) *ffapi.FilterResult); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*ffapi.FilterResult)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, ffapi.AndFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Name provides a mock function with given fields:
func (_m *Manager) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewMessage provides a mock function with given fields: msg
func (_m *Manager) NewMessage(msg *core.MessageInOut) syncasync.Sender {
	ret := _m.Called(msg)

	var r0 syncasync.Sender
	if rf, ok := ret.Get(0).(func(*core.MessageInOut) syncasync.Sender); ok {
		r0 = rf(msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(syncasync.Sender)
		}
	}

	return r0
}

// PrepareOperation provides a mock function with given fields: ctx, op
func (_m *Manager) PrepareOperation(ctx context.Context, op *core.Operation) (*core.PreparedOperation, error) {
	ret := _m.Called(ctx, op)

	var r0 *core.PreparedOperation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) (*core.PreparedOperation, error)); ok {
		return rf(ctx, op)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) *core.PreparedOperation); ok {
		r0 = rf(ctx, op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.PreparedOperation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Operation) error); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestReply provides a mock function with given fields: ctx, request
func (_m *Manager) RequestReply(ctx context.Context, request *core.MessageInOut) (*core.MessageInOut, error) {
	ret := _m.Called(ctx, request)

	var r0 *core.MessageInOut
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.MessageInOut) (*core.MessageInOut, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.MessageInOut) *core.MessageInOut); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.MessageInOut)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.MessageInOut) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveInitGroup provides a mock function with given fields: ctx, msg, creator
func (_m *Manager) ResolveInitGroup(ctx context.Context, msg *core.Message, creator *core.Member) (*core.Group, error) {
	ret := _m.Called(ctx, msg, creator)

	var r0 *core.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Message, *core.Member) (*core.Group, error)); ok {
		return rf(ctx, msg, creator)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Message, *core.Member) *core.Group); ok {
		r0 = rf(ctx, msg, creator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Message, *core.Member) error); ok {
		r1 = rf(ctx, msg, creator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunOperation provides a mock function with given fields: ctx, op
func (_m *Manager) RunOperation(ctx context.Context, op *core.PreparedOperation) (fftypes.JSONObject, bool, error) {
	ret := _m.Called(ctx, op)

	var r0 fftypes.JSONObject
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.PreparedOperation) (fftypes.JSONObject, bool, error)); ok {
		return rf(ctx, op)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.PreparedOperation) fftypes.JSONObject); ok {
		r0 = rf(ctx, op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fftypes.JSONObject)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.PreparedOperation) bool); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *core.PreparedOperation) error); ok {
		r2 = rf(ctx, op)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SendMessage provides a mock function with given fields: ctx, in, waitConfirm
func (_m *Manager) SendMessage(ctx context.Context, in *core.MessageInOut, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, in, waitConfirm)

	var r0 *core.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.MessageInOut, bool) (*core.Message, error)); ok {
		return rf(ctx, in, waitConfirm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.MessageInOut, bool) *core.Message); ok {
		r0 = rf(ctx, in, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.MessageInOut, bool) error); ok {
		r1 = rf(ctx, in, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewManager(t mockConstructorTestingTNewManager) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
