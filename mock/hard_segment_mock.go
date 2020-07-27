// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/listen-lavender/goseq/model (interfaces: HardSegmentDao)

// Package mock is a generated GoMock package.
package mock

import (
	model "github.com/listen-lavender/goseq/model"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHardSegmentDao is a mock of HardSegmentDao interface
type MockHardSegmentDao struct {
	ctrl     *gomock.Controller
	recorder *MockHardSegmentDaoMockRecorder
}

// MockHardSegmentDaoMockRecorder is the mock recorder for MockHardSegmentDao
type MockHardSegmentDaoMockRecorder struct {
	mock *MockHardSegmentDao
}

// NewMockHardSegmentDao creates a new mock instance
func NewMockHardSegmentDao(ctrl *gomock.Controller) *MockHardSegmentDao {
	mock := &MockHardSegmentDao{ctrl: ctrl}
	mock.recorder = &MockHardSegmentDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHardSegmentDao) EXPECT() *MockHardSegmentDaoMockRecorder {
	return m.recorder
}

// AtomicAdd mocks base method
func (m *MockHardSegmentDao) AtomicAdd(arg0 context.Context, arg1 *model.HardSegment) (*model.HardSegment, error) {
	ret := m.ctrl.Call(m, "AtomicAdd", arg0, arg1)
	ret0, _ := ret[0].(*model.HardSegment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtomicAdd indicates an expected call of AtomicAdd
func (mr *MockHardSegmentDaoMockRecorder) AtomicAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtomicAdd", reflect.TypeOf((*MockHardSegmentDao)(nil).AtomicAdd), arg0, arg1)
}

// AtomicUpdate mocks base method
func (m *MockHardSegmentDao) AtomicUpdate(arg0 context.Context, arg1 string, arg2 uint16, arg3 *model.HardSegment) (*model.HardSegment, error) {
	ret := m.ctrl.Call(m, "AtomicUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*model.HardSegment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AtomicUpdate indicates an expected call of AtomicUpdate
func (mr *MockHardSegmentDaoMockRecorder) AtomicUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtomicUpdate", reflect.TypeOf((*MockHardSegmentDao)(nil).AtomicUpdate), arg0, arg1, arg2, arg3)
}

// Find mocks base method
func (m *MockHardSegmentDao) Find(arg0 context.Context, arg1 uint64, arg2 int, arg3 string, arg4 []*model.HardSegment, arg5 func(*model.HardSegment) bool) ([]*model.HardSegment, error) {
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*model.HardSegment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockHardSegmentDaoMockRecorder) Find(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockHardSegmentDao)(nil).Find), arg0, arg1, arg2, arg3, arg4, arg5)
}

// StoreID mocks base method
func (m *MockHardSegmentDao) StoreID() string {
	ret := m.ctrl.Call(m, "StoreID")
	ret0, _ := ret[0].(string)
	return ret0
}

// StoreID indicates an expected call of StoreID
func (mr *MockHardSegmentDaoMockRecorder) StoreID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreID", reflect.TypeOf((*MockHardSegmentDao)(nil).StoreID))
}