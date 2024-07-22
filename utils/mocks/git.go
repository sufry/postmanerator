package mocks_test

import (
	. "github.com/sufry/postmanerator/utils"

	"github.com/stretchr/testify/mock"
)

type MockGitAgent struct {
	mock.Mock
}

func (m *MockGitAgent) Clone(args []string, options CloneOptions) error {
	return m.Called(args, options).Error(0)
}
