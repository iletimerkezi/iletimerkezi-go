package responses

type AccountResponse struct {
    *BaseResponse
    amount string
    sms    string
}

func NewAccountResponse(resp *Response) *AccountResponse {
    base := NewBaseResponse(resp)
    r := &AccountResponse{
        BaseResponse: base,
    }

    if balance, ok := r.Data["balance"].(map[string]interface{}); ok {
        if amount, ok := balance["amount"].(string); ok {
            r.amount = amount
        }
        if sms, ok := balance["sms"].(string); ok {
            r.sms = sms
        }
    }

    return r
}

func (r *AccountResponse) Amount() string {
    return r.amount
}

func (r *AccountResponse) Credits() string {
    return r.sms
} 