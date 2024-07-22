package mocks_test

import (
	"io"

	"github.com/stretchr/testify/mock"
	"github.com/sufry/postmanerator/postman"
	. "github.com/sufry/postmanerator/themes"
)

type MockThemeRenderer struct {
	mock.Mock
}

func (m *MockThemeRenderer) Render(w io.Writer, theme *Theme, collection postman.Collection) error {
	return m.Called(w, theme, collection).Error(0)
}
