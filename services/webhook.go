package services

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "github.com/iletimerkezi/iletimerkezi-go/models"
)

type WebhookService struct{}

func NewWebhookService() *WebhookService {
    return &WebhookService{}
}

func (s *WebhookService) Handle(r *http.Request) (*models.WebhookReport, error) {
    if r.Body == nil {
        return nil, fmt.Errorf("request body is empty")
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read request body: %w", err)
    }
    defer r.Body.Close()

    var data map[string]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, fmt.Errorf("failed to parse webhook data: %w", err)
    }

    return models.NewWebhookReport(data), nil
} 