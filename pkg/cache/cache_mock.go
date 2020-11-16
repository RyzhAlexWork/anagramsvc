package cache

import (
	"github.com/stretchr/testify/mock"
)

// MockCache ...
type MockCache struct {
	mock.Mock
}

// Get ...
func (m *MockCache) Get() (data []string) {
	args := m.Called()
	if a, ok := args.Get(0).([]string); ok {
		return a
	}
	return data
}

// Set ...
func (m *MockCache) Set(data []string) {
}
