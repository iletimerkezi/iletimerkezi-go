package services_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
)

func TestReportService(t *testing.T) {
    // Sample data based on PHP test
    sampleData := map[string]interface{}{
        "response": map[string]interface{}{
            "status": map[string]interface{}{
                "message": "Success",
            },
            "order": map[string]interface{}{
                "id":     "12345",
                "status": "114",
                "message": []interface{}{
                    map[string]interface{}{
                        "number": "905301234567",
                        "status": "111",
                    },
                    map[string]interface{}{
                        "number": "905301234568",
                        "status": "112",
                    },
                },
            },
        },
    }

    t.Run("get report successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := NewReportService(mockClient, "test-key", "test-hash")

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "order": map[string]interface{}{
                    "id":       12345,
                    "page":     1,
                    "rowCount": 1000,
                },
            },
        }

        mockClient.On("Post", "get-report/json", expectedPayload).
            Return(&Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.Get(12345, 1, 1000)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.Equal(t, "12345", resp.GetOrderID())
        assert.Equal(t, "COMPLETED", resp.GetOrderStatus())
        assert.True(t, resp.Ok())
    })

    t.Run("next page without previous request", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := NewReportService(mockClient, "test-key", "test-hash")

        resp, err := service.Next()
        assert.Error(t, err)
        assert.Nil(t, resp)
        assert.Contains(t, err.Error(), "no previous report request found")
    })
} 