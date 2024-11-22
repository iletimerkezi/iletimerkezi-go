package responses

type SmsResponse struct {
    *BaseResponse
    orderID string
}

func NewSmsResponse(resp *Response) *SmsResponse {
    base := NewBaseResponse(resp)
    
    r := &SmsResponse{
        BaseResponse: base,
    }
    
    if order, ok := r.Data["order"].(map[string]interface{}); ok {
        if id, ok := order["id"].(string); ok {
            r.orderID = id
        }
    }
    
    return r
}

func (r *SmsResponse) GetOrderID() string {
    return r.orderID
} 