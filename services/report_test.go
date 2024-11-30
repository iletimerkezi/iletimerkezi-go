package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
    "github.com/iletimerkezi/iletimerkezi-go/services"
)

func TestReportService(t *testing.T) {
    t.Run("get report successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewReportService(mockClient, "test-key", "test-hash")

        sampleData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    int(200),
                    "message": "İşlem başarılı",
                },
                "order": map[string]interface{}{
                    "id":          int(45),
                    "status":      int(113),
                    "total":       int(4),
                    "delivered":   int(0),
                    "undelivered": int(0),
                    "waiting":     int(4),
                    "submitAt":    "2024-10-12 23:27:42",
                    "sendAt":      "2024-10-12 23:27:42",
                    "sender":      "SENDER81",
                    "message": []interface{}{
                        map[string]interface{}{
                            "number": "+905057023100",
                            "status": int(110),
                        },
                        map[string]interface{}{
                            "number": "+905057023100",
                            "status": int(111),
                        },
                        map[string]interface{}{
                            "number": "+905057023100",
                            "status": int(112),
                        },
                        map[string]interface{}{
                            "number": "+905057023101",
                            "status": int(110),
                        },
                    },
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
                    "id":       45,
                    "page":     1,
                    "rowCount": 1000,
                },
            },
        }

        mockClient.On("Post", "get-report/json", expectedPayload).
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.Get(45, 1, 1000)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        
        // Check order details
        assert.Equal(t, 45, resp.OrderID)
        assert.Equal(t, "SENDING", resp.OrderStatus)
        assert.Equal(t, 113, resp.StatusCode)
        assert.Equal(t, 4, resp.Total)
        assert.Equal(t, 0, resp.Delivered)
        assert.Equal(t, 0, resp.Undelivered)
        assert.Equal(t, 4, resp.Waiting)
        assert.Equal(t, "2024-10-12 23:27:42", resp.SubmitAt)
        assert.Equal(t, "2024-10-12 23:27:42", resp.SendAt)
        assert.Equal(t, "SENDER81", resp.Sender)

        // Check messages
        messages := resp.Messages
        assert.Len(t, messages, 4)
        
        // Check first message
        assert.Equal(t, "+905057023100", messages[0].Number)
        assert.Equal(t, "WAITING", messages[0].Status)
        assert.Equal(t, 110, messages[0].StatusCode)
    })

    t.Run("get report with empty messages", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewReportService(mockClient, "test-key", "test-hash")

        emptyData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    int(200),
                    "message": "İşlem başarılı",
                },
                "order": map[string]interface{}{
                    "id":          45,
                    "status":      113,
                    "total":       0,
                    "delivered":   0,
                    "undelivered": 0,
                    "waiting":     0,
                    "submitAt":    "2024-10-12 23:27:42",
                    "sendAt":      "2024-10-12 23:27:42",
                    "sender":      "SENDER81",
                    "price":       "0,0000",
                    "message":     []interface{}{},
                },
            },
        }

        mockClient.On("Post", "get-report/json", map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "order": map[string]interface{}{
                    "id":       45,
                    "page":     1,
                    "rowCount": 1000,
                },
            },
        }).Return(&responses.Response{StatusCode: 200, Body: emptyData}, nil)

        resp, err := service.Get(45, 1, 1000)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Empty(t, resp.Messages)
        assert.Equal(t, 0, resp.Total)
    })

    t.Run("next page without previous request", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewReportService(mockClient, "test-key", "test-hash")

        resp, err := service.Next()
        assert.Error(t, err)
        assert.Nil(t, resp)
        assert.Contains(t, err.Error(), "no previous report request found")
    })

    t.Run("next page after successful request", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewReportService(mockClient, "test-key", "test-hash")

        // First call to Get()
        mockClient.On("Post", "get-report/json", map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "order": map[string]interface{}{
                    "id":       45,
                    "page":     1,
                    "rowCount": 1000,
                },
            },
        }).Return(&responses.Response{StatusCode: 200, Body: map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    200,
                    "message": "İşlem başarılı",
                },
                "order": map[string]interface{}{},
            },
        }}, nil)

        // Call Get first
        _, err := service.Get(45, 1, 1000)
        assert.NoError(t, err)

        // Then setup mock for Next()
        mockClient.On("Post", "get-report/json", map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "order": map[string]interface{}{
                    "id":       45,
                    "page":     2,
                    "rowCount": 1000,
                },
            },
        }).Return(&responses.Response{StatusCode: 200, Body: map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    200,
                    "message": "İşlem başarılı",
                },
                "order": map[string]interface{}{},
            },
        }}, nil)

        // Call Next
        resp, err := service.Next()
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
    })
} 