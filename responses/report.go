package responses

type ReportResponse struct {
    *BaseResponse
    messages []Message
    position int
}

type Message struct {
    Number string
    Status string
    Code   int
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
    orderStatusMap = map[string]string{
        "113": OrderStatusSending,
        "114": OrderStatusCompleted,
        "115": OrderStatusCanceled,
    }

    messageStatusMap = map[string]string{
        "110": MessageStatusWaiting,
        "111": MessageStatusDelivered,
        "112": MessageStatusUndelivered,
    }
)

func NewReportResponse(resp *Response) *ReportResponse {
    base := NewBaseResponse(resp)
    r := &ReportResponse{
        BaseResponse: base,
    }
    r.parseMessages()
    return r
}

func (r *ReportResponse) parseMessages() {
    if order, ok := r.Data["order"].(map[string]interface{}); ok {
        if msgs, ok := order["message"].([]interface{}); ok {
            r.messages = make([]Message, len(msgs))
            for i, msg := range msgs {
                if m, ok := msg.(map[string]interface{}); ok {
                    r.messages[i] = Message{
                        Number: m["number"].(string),
                        Status: messageStatusMap[m["status"].(string)],
                        Code:   int(m["status"].(float64)),
                    }
                }
            }
        }
    }
}

func (r *ReportResponse) GetOrderID() string {
    if order, ok := r.Data["order"].(map[string]interface{}); ok {
        return order["id"].(string)
    }
    return ""
}

func (r *ReportResponse) GetOrderStatus() string {
    if order, ok := r.Data["order"].(map[string]interface{}); ok {
        if status, ok := order["status"].(string); ok {
            return orderStatusMap[status]
        }
    }
    return ""
}

func (r *ReportResponse) GetMessages() []Message {
    return r.messages
}

// Iterator implementation
func (r *ReportResponse) Next() bool {
    r.position++
    return r.position < len(r.messages)
}

func (r *ReportResponse) Current() *Message {
    if r.position >= len(r.messages) {
        return nil
    }
    return &r.messages[r.position]
}

func (r *ReportResponse) Reset() {
    r.position = 0
} 