// Code generated by mockery v2.20.2. DO NOT EDIT.

package apiservermocks

import (
	context "context"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"
	core "github.com/hyperledger/firefly/pkg/core"

	mock "github.com/stretchr/testify/mock"

	openapi3 "github.com/getkin/kin-openapi/openapi3"
)

// FFISwaggerGen is an autogenerated mock type for the FFISwaggerGen type
type FFISwaggerGen struct {
	mock.Mock
}

// Generate provides a mock function with given fields: ctx, baseURL, api, ffi
func (_m *FFISwaggerGen) Generate(ctx context.Context, baseURL string, api *core.ContractAPI, ffi *fftypes.FFI) *openapi3.T {
	ret := _m.Called(ctx, baseURL, api, ffi)

	var r0 *openapi3.T
	if rf, ok := ret.Get(0).(func(context.Context, string, *core.ContractAPI, *fftypes.FFI) *openapi3.T); ok {
		r0 = rf(ctx, baseURL, api, ffi)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openapi3.T)
		}
	}

	return r0
}

type mockConstructorTestingTNewFFISwaggerGen interface {
	mock.TestingT
	Cleanup(func())
}

// NewFFISwaggerGen creates a new instance of FFISwaggerGen. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFFISwaggerGen(t mockConstructorTestingTNewFFISwaggerGen) *FFISwaggerGen {
	mock := &FFISwaggerGen{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
