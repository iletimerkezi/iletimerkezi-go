package iletimerkezi

import (
	"github.com/iletimerkezi/iletimerkezi-go/services"
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

func (c *IletiMerkeziClient) SMS() *services.SmsService {
	return services.NewSmsService(c.httpClient, c.apiKey, c.apiHash, c.defaultSender)
}

func (c *IletiMerkeziClient) Reports() *services.ReportService {
	return services.NewReportService(c.httpClient, c.apiKey, c.apiHash)
}

func (c *IletiMerkeziClient) Summary() *services.SummaryService {
	return services.NewSummaryService(c.httpClient, c.apiKey, c.apiHash)
}

func (c *IletiMerkeziClient) Senders() *services.SenderService {
	return services.NewSenderService(c.httpClient, c.apiKey, c.apiHash)
}

func (c *IletiMerkeziClient) Account() *services.AccountService {
	return services.NewAccountService(c.httpClient, c.apiKey, c.apiHash)
}

func (c *IletiMerkeziClient) Webhook() *services.WebhookService {
	return services.NewWebhookService()
} 