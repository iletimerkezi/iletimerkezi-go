package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
    "github.com/iletimerkezi/iletimerkezi-go/services"
)

func TestSmsService(t *testing.T) {
    t.Run("send single message successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewSmsService(mockClient, "test-key", "test-hash", "SENDER")

        sampleData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    200,
                    "message": "İşlem başarılı",
                },
                "order": map[string]interface{}{
                    "id": 12323232,
                },
            },
        }

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "order": map[string]interface{}{
                    "sender": "SENDER",
                    "sendDateTime": "",
                    "iys": "1",
                    "iysList": "BIREYSEL",
                    "message": map[string]interface{}{
                        "text": "Test message",
                        "receipents": map[string][]string{
                            "number": []string{"5551234567"},
                        },
                    },
                },
            },
        }

        mockClient.On("Post", "send-sms/json", expectedPayload).
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.Send("5551234567", "Test message", "")
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, 12323232, resp.OrderID)
    })

    t.Run("send multiple messages successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewSmsService(mockClient, "test-key", "test-hash", "SENDER")

        sampleData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    200,
                    "message": "İşlem başarılı",
                },
                "order": map[string]interface{}{
                    "id": 12323232,
                },
            },
        }

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "order": map[string]interface{}{
                    "sender": "SENDER",
                    "sendDateTime": "",
                    "iys": "1",
                    "iysList": "BIREYSEL",
                    "message": map[string]interface{}{
                        "text": "Test message",
                        "receipents": map[string][]string{
                            "number": []string{
                                "5551234567",
                                "5551234568",
                            },
                        },
                    },
                },
            },
        }

        mockClient.On("Post", "send-sms/json", expectedPayload).
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.Send([]string{"5551234567", "5551234568"}, "Test message", "")
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, 12323232, resp.OrderID)
    })

    t.Run("send message with invalid credentials", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewSmsService(mockClient, "test-key", "test-hash", "SENDER")

        errorData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    401,
                    "message": "Üyelik bilgileri hatalı",
                },
            },
        }

		expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "order": map[string]interface{}{
                    "sender": "SENDER",
                    "sendDateTime": "",
                    "iys": "1",
                    "iysList": "BIREYSEL",
                    "message": map[string]interface{}{
                        "text": "Test message",
                        "receipents": map[string][]string{
                            "number": []string{
                                "5551234567",
                            },
                        },
                    },
                },
            },
        }

        mockClient.On("Post", "send-sms/json", expectedPayload).
            Return(&responses.Response{StatusCode: 401, Body: errorData}, nil)

        resp, err := service.Send("5551234567", "Test message", "")
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.False(t, resp.Ok())
        assert.Equal(t, 401, resp.GetStatusCode())
        assert.Equal(t, "Üyelik bilgileri hatalı", resp.GetMessage())
    })
}