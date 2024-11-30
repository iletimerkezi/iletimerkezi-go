package testutil

import (
	"github.com/stretchr/testify/mock"
	"github.com/iletimerkezi/iletimerkezi-go/responses"
)

type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Post(url string, payload interface{}) (*responses.Response, error) {
	args := m.Called(url, payload)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*responses.Response), args.Error(1)
}

func (m *MockHttpClient) GetLastResponse() *responses.Response {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*responses.Response)
}

func (m *MockHttpClient) GetLastPayload() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
} 