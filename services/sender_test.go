package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
)

func TestSenderService(t *testing.T) {
    sampleData := map[string]interface{}{
        "response": map[string]interface{}{
            "status": map[string]interface{}{
                "message": "Success",
            },
            "senders": map[string]interface{}{
                "sender": []interface{}{
                    "COMPANY1",
                    "COMPANY2",
                },
            },
        },
    }

    t.Run("list senders successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := NewSenderService(mockClient, "test-key", "test-hash")

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
            },
        }

        mockClient.On("Post", "get-sender/json", expectedPayload).
            Return(&Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.List()
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, []string{"COMPANY1", "COMPANY2"}, resp.GetSenders())
    })
} 