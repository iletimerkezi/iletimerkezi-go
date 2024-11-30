package responses

type Order struct {
    ID          int    `json:"id"`
    Status      int    `json:"status"`
    StatusText  string
    Total       int    `json:"total"`
    Delivered   int    `json:"delivered"`
    Undelivered int    `json:"undelivered"`
    Waiting     int    `json:"waiting"`
    SubmitAt    string `json:"submitAt"`
    SendAt      string `json:"sendAt"`
    Sender      string `json:"sender"`
}

type SummaryResponse struct {
    *BaseResponse
    Count  int     `json:"count"`
    Orders []Order `json:"orders"`
}

const (
    StatusSending   = 113
    StatusCompleted = 114
    StatusCanceled  = 115
)

func NewSummaryResponse(resp *Response) *SummaryResponse {
    base := NewBaseResponse(resp)
    r := &SummaryResponse{
        BaseResponse: base,
    }
    r.parseOrders()
    return r
}

func (r *SummaryResponse) parseOrders() {
    if r.Data == nil {
        return
    }

    // Parse count
    if count, ok := r.Data["count"].(float64); ok {
        r.Count = int(count)
    }

    // Parse orders
    if orders, ok := r.Data["orders"].([]interface{}); ok {
        r.Orders = make([]Order, 0, len(orders))
        for _, o := range orders {
            if orderMap, ok := o.(map[string]interface{}); ok {
                order := Order{}
                
                // Parse integer fields
                if id, ok := orderMap["id"].(int); ok {
                    order.ID = id
                }
                if status, ok := orderMap["status"].(int); ok {
                    order.Status = status
                    order.StatusText = GetStatusText(order.Status)
                }
                if total, ok := orderMap["total"].(int); ok {
                    order.Total = total
                }
                if delivered, ok := orderMap["delivered"].(int); ok {
                    order.Delivered = delivered
                }
                if undelivered, ok := orderMap["undelivered"].(int); ok {
                    order.Undelivered = undelivered
                }
                if waiting, ok := orderMap["waiting"].(int); ok {
                    order.Waiting = waiting
                }

                // Parse string fields
                if submitAt, ok := orderMap["submitAt"].(string); ok {
                    order.SubmitAt = submitAt
                }
                if sendAt, ok := orderMap["sendAt"].(string); ok {
                    order.SendAt = sendAt
                }
                if sender, ok := orderMap["sender"].(string); ok {
                    order.Sender = sender
                }

                r.Orders = append(r.Orders, order)
            }
        }
    }
}

// Order helper methods
func GetStatusText(status int) string {
    switch status {
    case StatusSending:
        return "SENDING"
    case StatusCompleted:
        return "COMPLETED"
    case StatusCanceled:
        return "CANCELED"
    default:
        return "UNKNOWN"
    }
}