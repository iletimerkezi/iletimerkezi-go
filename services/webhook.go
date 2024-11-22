package services

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type WebhookService struct{}

func NewWebhookService() *WebhookService {
    return &WebhookService{}
}

func (s *WebhookService) Handle(r *http.Request) (*WebhookReport, error) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()

    if len(body) == 0 {
        return nil, fmt.Errorf("no POST data received")
    }

    var data map[string]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, fmt.Errorf("invalid JSON payload: %w", err)
    }

    return NewWebhookReport(data), nil
} 