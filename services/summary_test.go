package services_test

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"  // Eksik import eklendi
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
    "github.com/iletimerkezi/iletimerkezi-go/services"
)

func TestSummaryService(t *testing.T) {
    t.Run("get summary successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewSummaryService(mockClient, "test-key", "test-hash")

        startDate := time.Date(2024, 11, 11, 0, 0, 0, 0, time.UTC)
        endDate := time.Date(2024, 11, 20, 0, 0, 0, 0, time.UTC)

        sampleData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    200,
                    "message": "İşlem başarılı",
                },
                "count": 2,
                "orders": []interface{}{
                    map[string]interface{}{
                        "id":          89,
                        "status":      113,
                        "total":       1,
                        "delivered":   0,
                        "undelivered": 0,
                        "waiting":     1,
                        "submitAt":    "2024-11-12 02:04:08",
                        "sendAt":      "2024-11-12 02:04:08",
                        "sender":      "eMarka",
                    },
                    map[string]interface{}{
                        "id":          98,
                        "status":      113,
                        "total":       1,
                        "delivered":   0,
                        "undelivered": 0,
                        "waiting":     1,
                        "submitAt":    "2024-11-13 19:04:08",
                        "sendAt":      "2024-11-13 19:04:08",
                        "sender":      "eMarka",
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
                "filter": map[string]interface{}{
                    "start": "2024-11-11",
                    "end":   "2024-11-20",
                    "page":  1,
                },
            },
        }

        mockClient.On("Post", "get-report-summary/json", expectedPayload).
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.Get(&startDate, &endDate, 1)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())

        orders := resp.Orders
        assert.Len(t, orders, 2)

        // İlk sipariş kontrolü
        assert.Equal(t, 89, orders[0].ID)
        assert.Equal(t, "SENDING", orders[0].StatusText)
        assert.Equal(t, "2024-11-12 02:04:08", orders[0].SubmitAt)
        assert.Equal(t, "eMarka", orders[0].Sender)

        // İkinci sipariş kontrolü
        assert.Equal(t, 98, orders[1].ID)
        assert.Equal(t, "SENDING", orders[1].StatusText)
        assert.Equal(t, "2024-11-13 19:04:08", orders[1].SubmitAt)
        assert.Equal(t, "eMarka", orders[1].Sender)
    })

    t.Run("get summary with invalid credentials", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewSummaryService(mockClient, "invalid-key", "invalid-hash")

        startDate := time.Date(2024, 11, 11, 0, 0, 0, 0, time.UTC)
        endDate := time.Date(2024, 11, 20, 0, 0, 0, 0, time.UTC)

        errorData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    401,
                    "message": "Üyelik bilgileri hatalı",
                },
            },
        }

        mockClient.On("Post", "get-report-summary/json", mock.Anything).
            Return(&responses.Response{StatusCode: 401, Body: errorData}, nil)

        resp, err := service.Get(&startDate, &endDate, 1)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.False(t, resp.Ok())
        assert.Equal(t, 401, resp.StatusCode)
        assert.Equal(t, "Üyelik bilgileri hatalı", resp.GetMessage())
        assert.Empty(t, resp.Orders)
    })

    t.Run("get summary with empty date range", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewSummaryService(mockClient, "test-key", "test-hash")

        sampleData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    200,
                    "message": "İşlem başarılı",
                },
                "count":  0,
                "orders": []interface{}{},
            },
        }

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "filter": map[string]interface{}{
                    "page": 1,
                },
            },
        }

        mockClient.On("Post", "get-report-summary/json", expectedPayload).
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.Get(nil, nil, 1)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Empty(t, resp.Orders)
    })
}