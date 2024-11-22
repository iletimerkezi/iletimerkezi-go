package responses

type BaseResponse struct {
    StatusCode int
    Message    string
    Data      map[string]interface{}
}

func NewBaseResponse(resp *Response) *BaseResponse {
    var message string
    if status, ok := resp.Body["response"].(map[string]interface{}); ok {
        if statusMsg, ok := status["status"].(map[string]interface{}); ok {
            if msg, ok := statusMsg["message"].(string); ok {
                message = msg
            }
        }
    }

    return &BaseResponse{
        StatusCode: resp.StatusCode,
        Message:    message,
        Data:      resp.Body["response"].(map[string]interface{}),
    }
}

func (r *BaseResponse) Ok() bool {
    return r.StatusCode == 200
}

func (r *BaseResponse) GetStatusCode() int {
    return r.StatusCode
}

func (r *BaseResponse) GetMessage() string {
    return r.Message
} 