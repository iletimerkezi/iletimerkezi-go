package services_test

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/iletimerkezi/iletimerkezi-go/internal/testutil"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
    "github.com/iletimerkezi/iletimerkezi-go/services"
)

func TestBlacklistService(t *testing.T) {
    sampleData := map[string]interface{}{
        "response": map[string]interface{}{
            "status": map[string]interface{}{
                "code":    int(200),
                "message": "İşlem başarılı",
            },
            "blacklist": map[string]interface{}{
                "count": 2,
                "number": []interface{}{
                    "+905301234567",
                    "+905301234568",
                },
            },
        },
    }

    t.Run("list blacklist successfully", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewBlacklistService(mockClient, "test-key", "test-hash")

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
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.List(&startDate, &endDate, 1, 1000)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, 2, resp.Total)
        assert.Equal(t, []string{"+905301234567", "+905301234568"}, resp.Numbers)
    })

    t.Run("list blacklist without date filters", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewBlacklistService(mockClient, "test-key", "test-hash")

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
            Return(&responses.Response{StatusCode: 200, Body: sampleData}, nil)

        resp, err := service.List(nil, nil, 1, 1000)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
    })

    t.Run("list empty blacklist", func(t *testing.T) {
        mockClient := new(testutil.MockHttpClient)
        service := services.NewBlacklistService(mockClient, "test-key", "test-hash")

        emptyData := map[string]interface{}{
            "response": map[string]interface{}{
                "status": map[string]interface{}{
                    "code":    int(200),
                    "message": "İşlem başarılı",
                },
                "blacklist": map[string]interface{}{
                    "count": 0,
                    "number": []interface{}{},
                },
            },
        }

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
            Return(&responses.Response{StatusCode: 200, Body: emptyData}, nil)

        resp, err := service.List(nil, nil, 1, 1000)
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        assert.True(t, resp.Ok())
        assert.Equal(t, 0, resp.Total)
        assert.Empty(t, resp.Numbers)
        assert.False(t, resp.HasMorePages)
    })
}