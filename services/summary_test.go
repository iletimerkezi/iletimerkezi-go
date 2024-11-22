package services_test

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/yourusername/iletimerkezi-go/internal/testutil"
)

func TestSummaryService(t *testing.T) {
    // Sample data based on PHP test structure
    sampleData := map[string]interface{}{
        "response": map[string]interface{}{
            "status": map[string]interface{}{
                "message": "Success",
            },
            "orders": []interface{}{
                map[string]interface{}{
                    "id":     "12345",
                    "status": "114",
                    "date":   "2024-03-27 10:00:00",
                },
            },
        },
    }

    t.Run("list reports successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := NewSummaryService(mockClient, "test-key", "test-hash")

        startDate := time.Date(2024, 3, 27, 0, 0, 0, 0, time.UTC)
        endDate := time.Date(2024, 3, 27, 23, 59, 59, 0, time.UTC)

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "filter": map[string]interface{}{
                    "start": startDate.Format("2006-01-02"),
                    "end":   endDate.Format("2006-01-02"),
                    "page":  1,
                },
            },
        }

        mockClient.On("Post", "get-reports/json", expectedPayload).
            Return(&Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.List(startDate, endDate, 1)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, 1, len(resp.GetOrders()))
    })

    t.Run("next page without previous request", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := NewSummaryService(mockClient, "test-key", "test-hash")

        resp, err := service.Next()
        assert.Error(t, err)
        assert.Nil(t, resp)
        assert.Contains(t, err.Error(), "no previous report request found")
    })
} 