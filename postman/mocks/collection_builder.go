package mocks_test

import (
	"github.com/stretchr/testify/mock"
	. "github.com/sufry/postmanerator/postman"
)

type MockCollectionBuilder struct {
	mock.Mock
}

func (m *MockCollectionBuilder) FromFile(file string, options BuilderOptions) (Collection, error) {
	args := m.Called(file, options)
	return args.Get(0).(Collection), args.Error(1)
}
