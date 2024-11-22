package iletimerkezi

import (
	"fmt"
	"net/http"
)

type IletiMerkeziClient struct {
	apiKey        string
	apiHash       string
	defaultSender string
	httpClient    HttpClient
	debug         bool
}

// NewClient creates a new IletiMerkezi client
func NewClient(apiKey, apiHash string, opts ...ClientOption) *IletiMerkeziClient {
	client := &IletiMerkeziClient{
		apiKey:     apiKey,
		apiHash:    apiHash,
		httpClient: NewDefaultHttpClient(),
	}

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// ClientOption defines the option pattern for client configuration
type ClientOption func(*IletiMerkeziClient)

// WithDefaultSender sets the default sender ID
func WithDefaultSender(sender string) ClientOption {
	return func(c *IletiMerkeziClient) {
		c.defaultSender = sender
	}
}

// WithHttpClient sets a custom HTTP client
func WithHttpClient(httpClient HttpClient) ClientOption {
	return func(c *IletiMerkeziClient) {
		c.httpClient = httpClient
	}
}

// WithDebug enables debug mode
func WithDebug(debug bool) ClientOption {
	return func(c *IletiMerkeziClient) {
		c.debug = debug
	}
} 