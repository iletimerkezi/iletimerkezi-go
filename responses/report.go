package responses

type ReportResponse struct {
    *BaseResponse
    OrderID      int
    OrderStatus  string
    StatusCode   int
    Total       int
    Delivered   int
    Undelivered int
    Waiting     int
    SubmitAt    string
    SendAt      string
    Sender      string
    Messages    []Message
}

type Message struct {
    Number     string
    Status     string
    StatusCode int
}

const (
    OrderStatusSending   = "SENDING"
    OrderStatusCompleted = "COMPLETED"
    OrderStatusCanceled  = "CANCELED"

    MessageStatusWaiting     = "WAITING"
    MessageStatusDelivered   = "DELIVERED"
    MessageStatusUndelivered = "UNDELIVERED"
)

var (
    orderStatusMap = map[int]string{
        113: OrderStatusSending,
        114: OrderStatusCompleted,
        115: OrderStatusCanceled,
    }

    messageStatusMap = map[int]string{
        110: MessageStatusWaiting,
        111: MessageStatusDelivered,
        112: MessageStatusUndelivered,
    }
)

func NewReportResponse(resp *Response) *ReportResponse {
    base := NewBaseResponse(resp)
    r := &ReportResponse{
        BaseResponse: base,
    }
    r.parseOrderData()
    return r
}

func (r *ReportResponse) parseOrderData() {
    if order, ok := r.Data["order"].(map[string]interface{}); ok {
        // Parse order ID
        if id, ok := order["id"].(int); ok {
            r.OrderID = id
        }

        // Parse status
        if status, ok := order["status"].(int); ok {
            r.StatusCode = status
            r.OrderStatus = orderStatusMap[r.StatusCode]
        }

        // Parse counts
        if total, ok := order["total"].(int); ok {
            r.Total = total
        }
        if delivered, ok := order["delivered"].(int); ok {
            r.Delivered = delivered
        }
        if undelivered, ok := order["undelivered"].(int); ok {
            r.Undelivered = undelivered
        }
        if waiting, ok := order["waiting"].(int); ok {
            r.Waiting = waiting
        }

        // Parse dates and sender
        if submitAt, ok := order["submitAt"].(string); ok {
            r.SubmitAt = submitAt
        }
        if sendAt, ok := order["sendAt"].(string); ok {
            r.SendAt = sendAt
        }
        if sender, ok := order["sender"].(string); ok {
            r.Sender = sender
        }

        // Parse messages
        if msgs, ok := order["message"].([]interface{}); ok {
            r.Messages = make([]Message, 0, len(msgs))
            for _, msg := range msgs {
                if m, ok := msg.(map[string]interface{}); ok {
                    var msgStatus int
                    if status, ok := m["status"].(int); ok {
                        msgStatus = status
                    }

                    message := Message{
                        Number:     m["number"].(string),
                        Status:     messageStatusMap[msgStatus],
                        StatusCode: msgStatus,
                    }
                    r.Messages = append(r.Messages, message)
                }
            }
        }
    }
}