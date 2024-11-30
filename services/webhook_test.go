package services_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/iletimerkezi/iletimerkezi-go/services"
)

func TestWebhookService(t *testing.T) {
    t.Run("parse delivered webhook successfully", func(t *testing.T) {
        service := services.NewWebhookService()

        webhookData := map[string]interface{}{
            "report": map[string]interface{}{
                "id":        int(10),
                "packet_id": int(20),
                "status":    "delivered",
                "to":        "+905057023100",
                "body":      "Test mesaji",
            },
        }

        // JSON verisini oluştur
        jsonData, err := json.Marshal(webhookData)
        assert.NoError(t, err)

        // HTTP request oluştur
        req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(jsonData))
        req.Header.Set("Content-Type", "application/json")

        report, err := service.Handle(req)

        assert.NoError(t, err)
        assert.NotNil(t, report)
        assert.Equal(t, 10, report.ID)
        assert.Equal(t, 20, report.PacketID)
        assert.Equal(t, "delivered", report.Status)
        assert.Equal(t, "+905057023100", report.To)
        assert.Equal(t, "Test mesaji", report.Body)
        assert.True(t, report.IsDelivered())
        assert.False(t, report.IsUndelivered())
        assert.False(t, report.IsAccepted())
    })

    t.Run("handle invalid json data", func(t *testing.T) {
        service := services.NewWebhookService()

        // Geçersiz JSON verisi
        invalidJSON := []byte(`{"invalid json`)

        req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(invalidJSON))
        req.Header.Set("Content-Type", "application/json")

        report, err := service.Handle(req)
        assert.Error(t, err)
        assert.Nil(t, report)
    })

    t.Run("handle empty request body", func(t *testing.T) {
        service := services.NewWebhookService()

        req := httptest.NewRequest(http.MethodPost, "/webhook", nil)
        req.Header.Set("Content-Type", "application/json")

        report, err := service.Handle(req)
        assert.Error(t, err)
        assert.Nil(t, report)
    })

    t.Run("handle request with missing fields", func(t *testing.T) {
        service := services.NewWebhookService()

        webhookData := map[string]interface{}{
            "report": map[string]interface{}{
                "id": 10,
                // Diğer alanlar eksik
            },
        }

        jsonData, err := json.Marshal(webhookData)
        assert.NoError(t, err)

        req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(jsonData))
        req.Header.Set("Content-Type", "application/json")

        report, err := service.Handle(req)
        assert.NoError(t, err)
        assert.NotNil(t, report)
        assert.Equal(t, 10, report.ID)
        assert.Empty(t, report.PacketID)
        assert.Empty(t, report.Status)
        assert.Empty(t, report.To)
        assert.Empty(t, report.Body)
    })
}