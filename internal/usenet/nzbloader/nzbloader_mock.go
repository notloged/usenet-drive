// Code generated by MockGen. DO NOT EDIT.
// Source: ./nzbloader.go

// Package nzbloader is a generated GoMock package.
package nzbloader

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nzb "github.com/javi11/usenet-drive/pkg/nzb"
	osfs "github.com/javi11/usenet-drive/pkg/osfs"
)

// MockNzbLoader is a mock of NzbLoader interface.
type MockNzbLoader struct {
	ctrl     *gomock.Controller
	recorder *MockNzbLoaderMockRecorder
}

// MockNzbLoaderMockRecorder is the mock recorder for MockNzbLoader.
type MockNzbLoaderMockRecorder struct {
	mock *MockNzbLoader
}

// NewMockNzbLoader creates a new mock instance.
func NewMockNzbLoader(ctrl *gomock.Controller) *MockNzbLoader {
	mock := &MockNzbLoader{ctrl: ctrl}
	mock.recorder = &MockNzbLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNzbLoader) EXPECT() *MockNzbLoaderMockRecorder {
	return m.recorder
}

// EvictFromCache mocks base method.
func (m *MockNzbLoader) EvictFromCache(name string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvictFromCache", name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// EvictFromCache indicates an expected call of EvictFromCache.
func (mr *MockNzbLoaderMockRecorder) EvictFromCache(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictFromCache", reflect.TypeOf((*MockNzbLoader)(nil).EvictFromCache), name)
}

// LoadFromFile mocks base method.
func (m *MockNzbLoader) LoadFromFile(name string) (*NzbCache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadFromFile", name)
	ret0, _ := ret[0].(*NzbCache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadFromFile indicates an expected call of LoadFromFile.
func (mr *MockNzbLoaderMockRecorder) LoadFromFile(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadFromFile", reflect.TypeOf((*MockNzbLoader)(nil).LoadFromFile), name)
}

// LoadFromFileReader mocks base method.
func (m *MockNzbLoader) LoadFromFileReader(f osfs.File) (*NzbCache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadFromFileReader", f)
	ret0, _ := ret[0].(*NzbCache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadFromFileReader indicates an expected call of LoadFromFileReader.
func (mr *MockNzbLoaderMockRecorder) LoadFromFileReader(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadFromFileReader", reflect.TypeOf((*MockNzbLoader)(nil).LoadFromFileReader), f)
}

// RefreshCachedNzb mocks base method.
func (m *MockNzbLoader) RefreshCachedNzb(name string, nzb *nzb.Nzb) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshCachedNzb", name, nzb)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshCachedNzb indicates an expected call of RefreshCachedNzb.
func (mr *MockNzbLoaderMockRecorder) RefreshCachedNzb(name, nzb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshCachedNzb", reflect.TypeOf((*MockNzbLoader)(nil).RefreshCachedNzb), name, nzb)
}
