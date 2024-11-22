package services_test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestWebhookService(t *testing.T) {
    t.Run("handle webhook data successfully", func(t *testing.T) {
        sampleData := `{
            "report": {
                "id": "12345",
                "packet_id": "67890",
                "status": "delivered",
                "to": "905301234567",
                "body": "Test message"
            }
        }`

        req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBufferString(sampleData))
        service := NewWebhookService()

        report, err := service.Handle(req)
        assert.NoError(t, err)
        assert.NotNil(t, report)
        assert.Equal(t, "12345", report.GetID())
        assert.Equal(t, "67890", report.GetPacketID())
        assert.Equal(t, "delivered", report.GetStatus())
        assert.Equal(t, "905301234567", report.GetTo())
        assert.Equal(t, "Test message", report.GetBody())
        assert.True(t, report.IsDelivered())
    })

    t.Run("handle empty post data", func(t *testing.T) {
        req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
        service := NewWebhookService()

        report, err := service.Handle(req)
        assert.Error(t, err)
        assert.Nil(t, report)
        assert.Contains(t, err.Error(), "no POST data received")
    })

    t.Run("handle invalid json data", func(t *testing.T) {
        req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBufferString("invalid json"))
        service := NewWebhookService()

        report, err := service.Handle(req)
        assert.Error(t, err)
        assert.Nil(t, report)
        assert.Contains(t, err.Error(), "invalid JSON payload")
    })
} 