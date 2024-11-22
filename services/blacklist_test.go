package services_test

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/yourusername/iletimerkezi-go/internal/testutil"
)

func TestBlacklistService(t *testing.T) {
    sampleData := map[string]interface{}{
        "response": map[string]interface{}{
            "status": map[string]interface{}{
                "message": "Success",
            },
            "blacklist": map[string]interface{}{
                "count": 2,
                "number": []interface{}{
                    "905301234567",
                    "905301234568",
                },
            },
        },
    }

    t.Run("list blacklist successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := NewBlacklistService(mockClient, "test-key", "test-hash")

        startDate := time.Date(2024, 3, 27, 1, 10, 0, 0, time.UTC)
        endDate := time.Date(2024, 3, 27, 23, 59, 59, 0, time.UTC)

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "blacklist": map[string]interface{}{
                    "page":     1,
                    "rowCount": 1000,
                    "filter": map[string]string{
                        "start": startDate.Format("2006-01-02 15:04:05"),
                        "end":   endDate.Format("2006-01-02 15:04:05"),
                    },
                },
            },
        }

        mockClient.On("Post", "get-blacklist/json", expectedPayload).
            Return(&Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.List(&startDate, &endDate, 1, 1000)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, 2, resp.GetCount())
        assert.Equal(t, []string{"905301234567", "905301234568"}, resp.GetNumbers())
    })

    t.Run("list blacklist without date filters", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := NewBlacklistService(mockClient, "test-key", "test-hash")

        expectedPayload := map[string]interface{}{
            "request": map[string]interface{}{
                "authentication": map[string]string{
                    "key":  "test-key",
                    "hash": "test-hash",
                },
                "blacklist": map[string]interface{}{
                    "page":     1,
                    "rowCount": 1000,
                },
            },
        }

        mockClient.On("Post", "get-blacklist/json", expectedPayload).
            Return(&Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.List(nil, nil, 1, 1000)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
    })
} 