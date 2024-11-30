package iletimerkezi

import (
	"github.com/iletimerkezi/iletimerkezi-go/services"
	"encoding/json"
	"fmt"
)

type IletiMerkeziClient struct {
	apiKey        string
	apiHash       string
	defaultSender string
	httpClient    HttpClient
}

type DebugInfo struct {
	Payload  interface{}     `json:"payload"`
	Response map[string]interface{} `json:"response"`
	Status   int            `json:"status"`
}

// NewClient creates a new IletiMerkezi client
func NewClient(apiKey, apiHash string) *IletiMerkeziClient {
	return NewClientWithHttpClient(apiKey, apiHash, NewHttpClient())
}

func NewClientWithHttpClient(apiKey, apiHash string, httpClient HttpClient) *IletiMerkeziClient {
	return &IletiMerkeziClient{
		apiKey:     apiKey,
		apiHash:    apiHash,
		httpClient: httpClient,
	}
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

func (c *IletiMerkeziClient) SetDefaultSender(sender string) {
	c.defaultSender = sender
}

func (c *IletiMerkeziClient) Sms() *services.SmsService {
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

func (c *IletiMerkeziClient) Blacklist() *services.BlacklistService {
	return services.NewBlacklistService(c.httpClient, c.apiKey, c.apiHash)
}

func (c *IletiMerkeziClient) Debug() string {

	lastResponse := c.httpClient.GetLastResponse()
	if lastResponse == nil {
		return "No request has been made yet"
	}

	debug := DebugInfo{
		Payload:  json.RawMessage(c.httpClient.GetLastPayload()),
		Response: lastResponse.Body,
		Status:   lastResponse.StatusCode,
	}

	debugJSON, err := json.MarshalIndent(debug, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error creating debug info: %v", err)
	}

	return string(debugJSON)
} 