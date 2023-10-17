// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/async (interfaces: Cache,TTLCache)

// Package mock_async is a generated GoMock package.
package mock_async

import (
	reflect "reflect"
	time "time"

	async "github.com/aws/amazon-ecs-agent/ecs-agent/async"
	gomock "github.com/golang/mock/gomock"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCache) Delete(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0)
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCache)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockCache) Get(arg0 string) (async.Value, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(async.Value)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockCache) Set(arg0 string, arg1 async.Value) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockCacheMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), arg0, arg1)
}

// MockTTLCache is a mock of TTLCache interface.
type MockTTLCache struct {
	ctrl     *gomock.Controller
	recorder *MockTTLCacheMockRecorder
}

// MockTTLCacheMockRecorder is the mock recorder for MockTTLCache.
type MockTTLCacheMockRecorder struct {
	mock *MockTTLCache
}

// NewMockTTLCache creates a new mock instance.
func NewMockTTLCache(ctrl *gomock.Controller) *MockTTLCache {
	mock := &MockTTLCache{ctrl: ctrl}
	mock.recorder = &MockTTLCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTTLCache) EXPECT() *MockTTLCacheMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockTTLCache) Delete(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0)
}

// Delete indicates an expected call of Delete.
func (mr *MockTTLCacheMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTTLCache)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockTTLCache) Get(arg0 string) (interface{}, bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockTTLCacheMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTTLCache)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockTTLCache) Set(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockTTLCacheMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockTTLCache)(nil).Set), arg0, arg1)
}

// SetTTL mocks base method.
func (m *MockTTLCache) SetTTL(arg0 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTTL", arg0)
}

// SetTTL indicates an expected call of SetTTL.
func (mr *MockTTLCacheMockRecorder) SetTTL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTTL", reflect.TypeOf((*MockTTLCache)(nil).SetTTL), arg0)
}