// Code generated by MockGen. DO NOT EDIT.
// Source: gachamachine/storage (interfaces: Uploader)

// Package mocks is a generated GoMock package.
package mocks

import (
	storage "gachamachine/storage"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockUploader is a mock of Uploader interface
type MockUploader struct {
	ctrl     *gomock.Controller
	recorder *MockUploaderMockRecorder
}

// MockUploaderMockRecorder is the mock recorder for MockUploader
type MockUploaderMockRecorder struct {
	mock *MockUploader
}

// NewMockUploader creates a new mock instance
func NewMockUploader(ctrl *gomock.Controller) *MockUploader {
	mock := &MockUploader{ctrl: ctrl}
	mock.recorder = &MockUploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUploader) EXPECT() *MockUploaderMockRecorder {
	return m.recorder
}

// Upload mocks base method
func (m *MockUploader) Upload(arg0 *storage.FileInfo, arg1 io.ReadCloser) error {
	ret := m.ctrl.Call(m, "Upload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockUploaderMockRecorder) Upload(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockUploader)(nil).Upload), arg0, arg1)
}
