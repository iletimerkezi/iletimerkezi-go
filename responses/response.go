package responses

type Response struct {
    StatusCode int
    Body       map[string]interface{}
}

func NewResponse(statusCode int, body map[string]interface{}) *Response {
    return &Response{
        StatusCode: statusCode,
        Body:       body,
    }
} 