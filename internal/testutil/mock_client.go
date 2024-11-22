package testutil

import (
	"github.com/stretchr/testify/mock"
)

type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Post(url string, payload interface{}) (*Response, error) {
	args := m.Called(url, payload)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Response), args.Error(1)
}

func (m *MockHttpClient) GetLastResponse() *Response {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*Response)
} 