package test

import "github.com/stretchr/testify/mock"

// Provide stub implementation of Terraform's ResourceData
// for unit testing

// NOTE: This only stubs out methods actually used

type StubResourceData struct {
	mock.Mock
}

func (m *StubResourceData) Get(key string) interface{} {
	args := m.Called(key)
	return args.Get(0)
}

func (m *StubResourceData) GetOk(key string) (interface{}, bool) {
	args := m.Called(key)
	return args.Get(0), args.Bool(1)
}
