package responses

type SmsResponse struct {
    *BaseResponse
    OrderID int
}

func NewSmsResponse(resp *Response) *SmsResponse {
    base := NewBaseResponse(resp)
    
    r := &SmsResponse{
        BaseResponse: base,
    }
    
    if order, ok := r.Data["order"].(map[string]interface{}); ok {
        if id, ok := order["id"].(int); ok {
            r.OrderID = id
        }
    }
    
    return r
}