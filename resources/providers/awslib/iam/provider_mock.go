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

	awslib "github.com/elastic/cloudbeat/resources/providers/awslib"

	mock "github.com/stretchr/testify/mock"
)

// MockAccessManagement is an autogenerated mock type for the AccessManagement type
type MockAccessManagement struct {
	mock.Mock
}

type MockAccessManagement_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccessManagement) EXPECT() *MockAccessManagement_Expecter {
	return &MockAccessManagement_Expecter{mock: &_m.Mock}
}

// GetIAMRolePermissions provides a mock function with given fields: ctx, roleName
func (_m *MockAccessManagement) GetIAMRolePermissions(ctx context.Context, roleName string) ([]RolePolicyInfo, error) {
	ret := _m.Called(ctx, roleName)

	var r0 []RolePolicyInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) []RolePolicyInfo); ok {
		r0 = rf(ctx, roleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RolePolicyInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, roleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_GetIAMRolePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetIAMRolePermissions'
type MockAccessManagement_GetIAMRolePermissions_Call struct {
	*mock.Call
}

// GetIAMRolePermissions is a helper method to define mock.On call
//  - ctx context.Context
//  - roleName string
func (_e *MockAccessManagement_Expecter) GetIAMRolePermissions(ctx interface{}, roleName interface{}) *MockAccessManagement_GetIAMRolePermissions_Call {
	return &MockAccessManagement_GetIAMRolePermissions_Call{Call: _e.mock.On("GetIAMRolePermissions", ctx, roleName)}
}

func (_c *MockAccessManagement_GetIAMRolePermissions_Call) Run(run func(ctx context.Context, roleName string)) *MockAccessManagement_GetIAMRolePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockAccessManagement_GetIAMRolePermissions_Call) Return(_a0 []RolePolicyInfo, _a1 error) *MockAccessManagement_GetIAMRolePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPasswordPolicy provides a mock function with given fields: ctx
func (_m *MockAccessManagement) GetPasswordPolicy(ctx context.Context) (awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	var r0 awslib.AwsResource
	if rf, ok := ret.Get(0).(func(context.Context) awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(awslib.AwsResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_GetPasswordPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPasswordPolicy'
type MockAccessManagement_GetPasswordPolicy_Call struct {
	*mock.Call
}

// GetPasswordPolicy is a helper method to define mock.On call
//  - ctx context.Context
func (_e *MockAccessManagement_Expecter) GetPasswordPolicy(ctx interface{}) *MockAccessManagement_GetPasswordPolicy_Call {
	return &MockAccessManagement_GetPasswordPolicy_Call{Call: _e.mock.On("GetPasswordPolicy", ctx)}
}

func (_c *MockAccessManagement_GetPasswordPolicy_Call) Run(run func(ctx context.Context)) *MockAccessManagement_GetPasswordPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccessManagement_GetPasswordPolicy_Call) Return(_a0 awslib.AwsResource, _a1 error) *MockAccessManagement_GetPasswordPolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUsers provides a mock function with given fields: ctx
func (_m *MockAccessManagement) GetUsers(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	var r0 []awslib.AwsResource
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccessManagement_GetUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsers'
type MockAccessManagement_GetUsers_Call struct {
	*mock.Call
}

// GetUsers is a helper method to define mock.On call
//  - ctx context.Context
func (_e *MockAccessManagement_Expecter) GetUsers(ctx interface{}) *MockAccessManagement_GetUsers_Call {
	return &MockAccessManagement_GetUsers_Call{Call: _e.mock.On("GetUsers", ctx)}
}

func (_c *MockAccessManagement_GetUsers_Call) Run(run func(ctx context.Context)) *MockAccessManagement_GetUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockAccessManagement_GetUsers_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockAccessManagement_GetUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockAccessManagement interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAccessManagement creates a new instance of MockAccessManagement. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAccessManagement(t mockConstructorTestingTNewMockAccessManagement) *MockAccessManagement {
	mock := &MockAccessManagement{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
