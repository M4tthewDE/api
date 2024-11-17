// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Chatterino/api/internal/db (interfaces: Pool)
//
// Generated by this command:
//
//	mockgen -destination ../mocks/mock_dbPool.go -package=mocks . Pool
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	pgconn "github.com/jackc/pgconn"
	v4 "github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	gomock "go.uber.org/mock/gomock"
)

// MockPool is a mock of Pool interface.
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
	isgomock struct{}
}

// MockPoolMockRecorder is the mock recorder for MockPool.
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance.
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// Acquire mocks base method.
func (m *MockPool) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Acquire", ctx)
	ret0, _ := ret[0].(*pgxpool.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Acquire indicates an expected call of Acquire.
func (mr *MockPoolMockRecorder) Acquire(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Acquire", reflect.TypeOf((*MockPool)(nil).Acquire), ctx)
}

// Begin mocks base method.
func (m *MockPool) Begin(arg0 context.Context) (v4.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", arg0)
	ret0, _ := ret[0].(v4.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockPoolMockRecorder) Begin(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockPool)(nil).Begin), arg0)
}

// Exec mocks base method.
func (m *MockPool) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, sql}
	for _, a := range arguments {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockPoolMockRecorder) Exec(ctx, sql any, arguments ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, sql}, arguments...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockPool)(nil).Exec), varargs...)
}

// Ping mocks base method.
func (m *MockPool) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockPoolMockRecorder) Ping(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockPool)(nil).Ping), ctx)
}

// Query mocks base method.
func (m *MockPool) Query(ctx context.Context, sql string, args ...any) (v4.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(v4.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockPoolMockRecorder) Query(ctx, sql any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockPool)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockPool) QueryRow(ctx context.Context, sql string, args ...any) v4.Row {
	m.ctrl.T.Helper()
	varargs := []any{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(v4.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockPoolMockRecorder) QueryRow(ctx, sql any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockPool)(nil).QueryRow), varargs...)
}
