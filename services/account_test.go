package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
    "github.com/iletimerkezi/iletimerkezi-go/services"
)

func TestAccountService(t *testing.T) {
    sampleData := map[string]interface{}{
        "response": map[string]interface{}{
            "status": map[string]interface{}{
                "code":    int(200),
                "message": "İşlem başarılı",
            },
            "balance": map[string]interface{}{
                "amount": float64(100.50),
                "sms":    int(1000),
            },
        },
    }

    t.Run("get balance successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewAccountService(mockClient, "test-key", "test-hash")

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
            },
        }

        mockClient.On("Post", "get-balance/json", expectedPayload).
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.Balance()
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, float64(100.50), resp.Amount)
        assert.Equal(t, int(1000), resp.Credits)
    })

    t.Run("get balance with invalid credentials", func(t *testing.T) {
        errorData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    int(401),
                    "message": "Üyelik bilgileri hatalı",
                },
            },
        }

        mockClient := new(testutil.MockHttpClient)
        service := services.NewAccountService(mockClient, "invalid-key", "invalid-hash")

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "invalid-key",
                    "hash": "invalid-hash",
                },
            },
        }

        mockClient.On("Post", "get-balance/json", expectedPayload).
            Return(&responses.Response{StatusCode: 401, Body: errorData}, nil)

        resp, err := service.Balance()
        assert.NoError(t, err)          // API çağrısı başarılı ama yanıt hata içeriyor
        assert.NotNil(t, resp)
        assert.False(t, resp.Ok())      // response.Ok() false olmalı
        assert.Equal(t, 401, resp.GetStatusCode())
        assert.Equal(t, "Üyelik bilgileri hatalı", resp.GetMessage())
        assert.Equal(t, float64(0), resp.Amount)    // Hata durumunda bakiye 0
        assert.Equal(t, 0, resp.Credits)            // Hata durumunda kredi 0
    })
} 