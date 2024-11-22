package responses_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestReportResponse(t *testing.T) {
    // Sample data based on PHP test (lines 13-33)
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

    t.Run("parses order data correctly", func(t *testing.T) {
        resp := NewReportResponse(&Response{StatusCode: 200, Body: sampleData})

        assert.Equal(t, "12345", resp.GetOrderID())
        assert.Equal(t, "COMPLETED", resp.GetOrderStatus())
        assert.Equal(t, 114, resp.GetOrderStatusCode())
        assert.True(t, resp.Ok())
    })

    t.Run("implements iterator correctly", func(t *testing.T) {
        resp := NewReportResponse(&Response{StatusCode: 200, Body: sampleData})
        messages := make([]map[string]string, 0)

        for resp.Next() {
            msg := resp.Current()
            messages = append(messages, map[string]string{
                "number": msg.Number,
                "status": msg.Status,
            })
        }

        assert.Len(t, messages, 2)
        assert.Equal(t, "905301234567", messages[0]["number"])
        assert.Equal(t, "905301234568", messages[1]["number"])
    })

    t.Run("translates message status correctly", func(t *testing.T) {
        resp := NewReportResponse(&Response{StatusCode: 200, Body: sampleData})
        
        assert.Equal(t, "DELIVERED", resp.GetMessageStatus(0))
        assert.Equal(t, "UNDELIVERED", resp.GetMessageStatus(1))
    })
} 