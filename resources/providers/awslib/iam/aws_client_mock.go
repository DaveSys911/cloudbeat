// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.13.1. DO NOT EDIT.

package iam

import (
	context "context"

	serviceiam "github.com/aws/aws-sdk-go-v2/service/iam"
	mock "github.com/stretchr/testify/mock"
)

// MockIAMClient is an autogenerated mock type for the Client type
type MockIAMClient struct {
	mock.Mock
}

type MockIAMClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIAMClient) EXPECT() *MockIAMClient_Expecter {
	return &MockIAMClient_Expecter{mock: &_m.Mock}
}

// GetAccessKeyLastUsed provides a mock function with given fields: ctx, params, optFns
func (_m *MockIAMClient) GetAccessKeyLastUsed(ctx context.Context, params *serviceiam.GetAccessKeyLastUsedInput, optFns ...func(*serviceiam.Options)) (*serviceiam.GetAccessKeyLastUsedOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *serviceiam.GetAccessKeyLastUsedOutput
	if rf, ok := ret.Get(0).(func(context.Context, *serviceiam.GetAccessKeyLastUsedInput, ...func(*serviceiam.Options)) *serviceiam.GetAccessKeyLastUsedOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceiam.GetAccessKeyLastUsedOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *serviceiam.GetAccessKeyLastUsedInput, ...func(*serviceiam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIAMClient_GetAccessKeyLastUsed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccessKeyLastUsed'
type MockIAMClient_GetAccessKeyLastUsed_Call struct {
	*mock.Call
}

// GetAccessKeyLastUsed is a helper method to define mock.On call
//  - ctx context.Context
//  - params *serviceiam.GetAccessKeyLastUsedInput
//  - optFns ...func(*serviceiam.Options)
func (_e *MockIAMClient_Expecter) GetAccessKeyLastUsed(ctx interface{}, params interface{}, optFns ...interface{}) *MockIAMClient_GetAccessKeyLastUsed_Call {
	return &MockIAMClient_GetAccessKeyLastUsed_Call{Call: _e.mock.On("GetAccessKeyLastUsed",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIAMClient_GetAccessKeyLastUsed_Call) Run(run func(ctx context.Context, params *serviceiam.GetAccessKeyLastUsedInput, optFns ...func(*serviceiam.Options))) *MockIAMClient_GetAccessKeyLastUsed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceiam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceiam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceiam.GetAccessKeyLastUsedInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIAMClient_GetAccessKeyLastUsed_Call) Return(_a0 *serviceiam.GetAccessKeyLastUsedOutput, _a1 error) *MockIAMClient_GetAccessKeyLastUsed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAccountPasswordPolicy provides a mock function with given fields: ctx, params, optFns
func (_m *MockIAMClient) GetAccountPasswordPolicy(ctx context.Context, params *serviceiam.GetAccountPasswordPolicyInput, optFns ...func(*serviceiam.Options)) (*serviceiam.GetAccountPasswordPolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *serviceiam.GetAccountPasswordPolicyOutput
	if rf, ok := ret.Get(0).(func(context.Context, *serviceiam.GetAccountPasswordPolicyInput, ...func(*serviceiam.Options)) *serviceiam.GetAccountPasswordPolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceiam.GetAccountPasswordPolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *serviceiam.GetAccountPasswordPolicyInput, ...func(*serviceiam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIAMClient_GetAccountPasswordPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountPasswordPolicy'
type MockIAMClient_GetAccountPasswordPolicy_Call struct {
	*mock.Call
}

// GetAccountPasswordPolicy is a helper method to define mock.On call
//  - ctx context.Context
//  - params *serviceiam.GetAccountPasswordPolicyInput
//  - optFns ...func(*serviceiam.Options)
func (_e *MockIAMClient_Expecter) GetAccountPasswordPolicy(ctx interface{}, params interface{}, optFns ...interface{}) *MockIAMClient_GetAccountPasswordPolicy_Call {
	return &MockIAMClient_GetAccountPasswordPolicy_Call{Call: _e.mock.On("GetAccountPasswordPolicy",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIAMClient_GetAccountPasswordPolicy_Call) Run(run func(ctx context.Context, params *serviceiam.GetAccountPasswordPolicyInput, optFns ...func(*serviceiam.Options))) *MockIAMClient_GetAccountPasswordPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceiam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceiam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceiam.GetAccountPasswordPolicyInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIAMClient_GetAccountPasswordPolicy_Call) Return(_a0 *serviceiam.GetAccountPasswordPolicyOutput, _a1 error) *MockIAMClient_GetAccountPasswordPolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetRolePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *MockIAMClient) GetRolePolicy(ctx context.Context, params *serviceiam.GetRolePolicyInput, optFns ...func(*serviceiam.Options)) (*serviceiam.GetRolePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *serviceiam.GetRolePolicyOutput
	if rf, ok := ret.Get(0).(func(context.Context, *serviceiam.GetRolePolicyInput, ...func(*serviceiam.Options)) *serviceiam.GetRolePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceiam.GetRolePolicyOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *serviceiam.GetRolePolicyInput, ...func(*serviceiam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIAMClient_GetRolePolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRolePolicy'
type MockIAMClient_GetRolePolicy_Call struct {
	*mock.Call
}

// GetRolePolicy is a helper method to define mock.On call
//  - ctx context.Context
//  - params *serviceiam.GetRolePolicyInput
//  - optFns ...func(*serviceiam.Options)
func (_e *MockIAMClient_Expecter) GetRolePolicy(ctx interface{}, params interface{}, optFns ...interface{}) *MockIAMClient_GetRolePolicy_Call {
	return &MockIAMClient_GetRolePolicy_Call{Call: _e.mock.On("GetRolePolicy",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIAMClient_GetRolePolicy_Call) Run(run func(ctx context.Context, params *serviceiam.GetRolePolicyInput, optFns ...func(*serviceiam.Options))) *MockIAMClient_GetRolePolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceiam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceiam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceiam.GetRolePolicyInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIAMClient_GetRolePolicy_Call) Return(_a0 *serviceiam.GetRolePolicyOutput, _a1 error) *MockIAMClient_GetRolePolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListAccessKeys provides a mock function with given fields: ctx, params, optFns
func (_m *MockIAMClient) ListAccessKeys(ctx context.Context, params *serviceiam.ListAccessKeysInput, optFns ...func(*serviceiam.Options)) (*serviceiam.ListAccessKeysOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *serviceiam.ListAccessKeysOutput
	if rf, ok := ret.Get(0).(func(context.Context, *serviceiam.ListAccessKeysInput, ...func(*serviceiam.Options)) *serviceiam.ListAccessKeysOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceiam.ListAccessKeysOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *serviceiam.ListAccessKeysInput, ...func(*serviceiam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIAMClient_ListAccessKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAccessKeys'
type MockIAMClient_ListAccessKeys_Call struct {
	*mock.Call
}

// ListAccessKeys is a helper method to define mock.On call
//  - ctx context.Context
//  - params *serviceiam.ListAccessKeysInput
//  - optFns ...func(*serviceiam.Options)
func (_e *MockIAMClient_Expecter) ListAccessKeys(ctx interface{}, params interface{}, optFns ...interface{}) *MockIAMClient_ListAccessKeys_Call {
	return &MockIAMClient_ListAccessKeys_Call{Call: _e.mock.On("ListAccessKeys",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIAMClient_ListAccessKeys_Call) Run(run func(ctx context.Context, params *serviceiam.ListAccessKeysInput, optFns ...func(*serviceiam.Options))) *MockIAMClient_ListAccessKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceiam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceiam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceiam.ListAccessKeysInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIAMClient_ListAccessKeys_Call) Return(_a0 *serviceiam.ListAccessKeysOutput, _a1 error) *MockIAMClient_ListAccessKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListAttachedRolePolicies provides a mock function with given fields: ctx, params, optFns
func (_m *MockIAMClient) ListAttachedRolePolicies(ctx context.Context, params *serviceiam.ListAttachedRolePoliciesInput, optFns ...func(*serviceiam.Options)) (*serviceiam.ListAttachedRolePoliciesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *serviceiam.ListAttachedRolePoliciesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *serviceiam.ListAttachedRolePoliciesInput, ...func(*serviceiam.Options)) *serviceiam.ListAttachedRolePoliciesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceiam.ListAttachedRolePoliciesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *serviceiam.ListAttachedRolePoliciesInput, ...func(*serviceiam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIAMClient_ListAttachedRolePolicies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAttachedRolePolicies'
type MockIAMClient_ListAttachedRolePolicies_Call struct {
	*mock.Call
}

// ListAttachedRolePolicies is a helper method to define mock.On call
//  - ctx context.Context
//  - params *serviceiam.ListAttachedRolePoliciesInput
//  - optFns ...func(*serviceiam.Options)
func (_e *MockIAMClient_Expecter) ListAttachedRolePolicies(ctx interface{}, params interface{}, optFns ...interface{}) *MockIAMClient_ListAttachedRolePolicies_Call {
	return &MockIAMClient_ListAttachedRolePolicies_Call{Call: _e.mock.On("ListAttachedRolePolicies",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIAMClient_ListAttachedRolePolicies_Call) Run(run func(ctx context.Context, params *serviceiam.ListAttachedRolePoliciesInput, optFns ...func(*serviceiam.Options))) *MockIAMClient_ListAttachedRolePolicies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceiam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceiam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceiam.ListAttachedRolePoliciesInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIAMClient_ListAttachedRolePolicies_Call) Return(_a0 *serviceiam.ListAttachedRolePoliciesOutput, _a1 error) *MockIAMClient_ListAttachedRolePolicies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListMFADevices provides a mock function with given fields: ctx, params, optFns
func (_m *MockIAMClient) ListMFADevices(ctx context.Context, params *serviceiam.ListMFADevicesInput, optFns ...func(*serviceiam.Options)) (*serviceiam.ListMFADevicesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *serviceiam.ListMFADevicesOutput
	if rf, ok := ret.Get(0).(func(context.Context, *serviceiam.ListMFADevicesInput, ...func(*serviceiam.Options)) *serviceiam.ListMFADevicesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceiam.ListMFADevicesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *serviceiam.ListMFADevicesInput, ...func(*serviceiam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIAMClient_ListMFADevices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListMFADevices'
type MockIAMClient_ListMFADevices_Call struct {
	*mock.Call
}

// ListMFADevices is a helper method to define mock.On call
//  - ctx context.Context
//  - params *serviceiam.ListMFADevicesInput
//  - optFns ...func(*serviceiam.Options)
func (_e *MockIAMClient_Expecter) ListMFADevices(ctx interface{}, params interface{}, optFns ...interface{}) *MockIAMClient_ListMFADevices_Call {
	return &MockIAMClient_ListMFADevices_Call{Call: _e.mock.On("ListMFADevices",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIAMClient_ListMFADevices_Call) Run(run func(ctx context.Context, params *serviceiam.ListMFADevicesInput, optFns ...func(*serviceiam.Options))) *MockIAMClient_ListMFADevices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceiam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceiam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceiam.ListMFADevicesInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIAMClient_ListMFADevices_Call) Return(_a0 *serviceiam.ListMFADevicesOutput, _a1 error) *MockIAMClient_ListMFADevices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListUsers provides a mock function with given fields: ctx, params, optFns
func (_m *MockIAMClient) ListUsers(ctx context.Context, params *serviceiam.ListUsersInput, optFns ...func(*serviceiam.Options)) (*serviceiam.ListUsersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *serviceiam.ListUsersOutput
	if rf, ok := ret.Get(0).(func(context.Context, *serviceiam.ListUsersInput, ...func(*serviceiam.Options)) *serviceiam.ListUsersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceiam.ListUsersOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *serviceiam.ListUsersInput, ...func(*serviceiam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIAMClient_ListUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListUsers'
type MockIAMClient_ListUsers_Call struct {
	*mock.Call
}

// ListUsers is a helper method to define mock.On call
//  - ctx context.Context
//  - params *serviceiam.ListUsersInput
//  - optFns ...func(*serviceiam.Options)
func (_e *MockIAMClient_Expecter) ListUsers(ctx interface{}, params interface{}, optFns ...interface{}) *MockIAMClient_ListUsers_Call {
	return &MockIAMClient_ListUsers_Call{Call: _e.mock.On("ListUsers",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIAMClient_ListUsers_Call) Run(run func(ctx context.Context, params *serviceiam.ListUsersInput, optFns ...func(*serviceiam.Options))) *MockIAMClient_ListUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*serviceiam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*serviceiam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*serviceiam.ListUsersInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIAMClient_ListUsers_Call) Return(_a0 *serviceiam.ListUsersOutput, _a1 error) *MockIAMClient_ListUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockIAMClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIAMClient creates a new instance of MockIAMClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIAMClient(t mockConstructorTestingTNewMockIAMClient) *MockIAMClient {
	mock := &MockIAMClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
