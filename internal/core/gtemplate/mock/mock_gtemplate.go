// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/gtemplate (interfaces: Repository)
//
// Generated by this command:
//
//	mockgen -destination=./mock/mock_gtemplate.go -package=mock github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/gtemplate Repository
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gtemplate "github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/gtemplate"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, obj *gtemplate.GameTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, obj)
}

// DB mocks base method.
func (m *MockRepository) DB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockRepositoryMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockRepository)(nil).DB))
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, obj *gtemplate.GameTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, obj)
}

// FindById mocks base method.
func (m *MockRepository) FindById(ctx context.Context, id int) (*gtemplate.GameTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, id)
	ret0, _ := ret[0].(*gtemplate.GameTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockRepositoryMockRecorder) FindById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepository)(nil).FindById), ctx, id)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, obj *gtemplate.GameTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, obj)
}
