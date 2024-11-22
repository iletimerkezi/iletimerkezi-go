package responses

type AccountResponse struct {
    response *Response
}

func NewAccountResponse(resp *Response) *AccountResponse {
    return &AccountResponse{
        response: resp,
    }
}

func (r *AccountResponse) Ok() bool {
    if status, ok := r.response.Body["response"].(map[string]interface{})["status"].(map[string]interface{}); ok {
        return status["message"] == "Success"
    }
    return false
}

func (r *AccountResponse) GetBalance() string {
    if balance, ok := r.response.Body["response"].(map[string]interface{})["balance"].(map[string]interface{}); ok {
        return balance["amount"].(string)
    }
    return "0"
}

func (r *AccountResponse) GetSmsCount() string {
    if balance, ok := r.response.Body["response"].(map[string]interface{})["balance"].(map[string]interface{}); ok {
        return balance["sms"].(string)
    }
    return "0"
} 