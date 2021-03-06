// Code generated by MockGen. DO NOT EDIT.
// Source: ./application/message/port/repository.go

// Package port is a generated GoMock package.
package port

import (
	domain "gin-gorm-sample/application/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(messageID uint) (domain.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", messageID)
	ret0, _ := ret[0].(domain.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), messageID)
}

// GetAll mocks base method
func (m *MockRepository) GetAll(userID uint) ([]domain.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userID)
	ret0, _ := ret[0].([]domain.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockRepositoryMockRecorder) GetAll(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll), userID)
}

// Create mocks base method
func (m *MockRepository) Create(message *domain.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), message)
}

// Update mocks base method
func (m *MockRepository) Update(message *domain.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), message)
}

// Delete mocks base method
func (m *MockRepository) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}
