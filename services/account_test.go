package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
)

func TestAccountService(t *testing.T) {
    sampleData := map[string]interface{}{
        "response": map[string]interface{}{
            "status": map[string]interface{}{
                "message": "Success",
            },
            "balance": map[string]interface{}{
                "amount": "100.50",
                "sms":    "1000",
            },
        },
    }

    t.Run("get balance successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := NewAccountService(mockClient, "test-key", "test-hash")

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
            },
        }

        mockClient.On("Post", "get-balance/json", expectedPayload).
            Return(&Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.Balance()
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, "100.50", resp.Amount())
        assert.Equal(t, "1000", resp.Credits())
    })
} 