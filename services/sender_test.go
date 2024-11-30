package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
    "github.com/iletimerkezi/iletimerkezi-go/services"
)

func TestSenderService(t *testing.T) {
    sampleData := map[string]interface{}{
        "response": map[string]interface{}{
            "status": map[string]interface{}{
                "code":    int(200),
                "message": "İşlem başarılı",
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
        service := services.NewSenderService(mockClient, "test-key", "test-hash")

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
            },
        }

        mockClient.On("Post", "get-sender/json", expectedPayload).
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.List()
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, []string{"COMPANY1", "COMPANY2"}, resp.Senders)
    })

    t.Run("list empty senders", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewSenderService(mockClient, "test-key", "test-hash")

        emptyData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    int(200),
                    "message": "İşlem başarılı",
                },
                "senders": map[string]interface{}{
                    "sender": []interface{}{},
                },
            },
        }

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
            },
        }

        mockClient.On("Post", "get-sender/json", expectedPayload).
            Return(&responses.Response{StatusCode: 200, Body: emptyData}, nil)

        resp, err := service.List()
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Empty(t, resp.Senders)
        assert.Len(t, resp.Senders, 0)
    })
} 