package responses

type AccountResponse struct {
    *BaseResponse
    Amount  float64
    Credits int
}

func NewAccountResponse(resp *Response) *AccountResponse {
    base := NewBaseResponse(resp)
    
    var amount float64 = 0
    var credits int = 0
    
    if response, ok := resp.Body["response"].(map[string]interface{}); ok {
        if balance, ok := response["balance"].(map[string]interface{}); ok {
            if amt, ok := balance["amount"].(float64); ok {
                amount = amt
            }
            if sms, ok := balance["sms"].(int); ok {
                credits = sms
            }
        }
    }
    
    return &AccountResponse{
        BaseResponse: base,
        Amount:      amount,
        Credits:     credits,
    }
}