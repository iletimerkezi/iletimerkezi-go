package services

type Response struct {
    StatusCode int
    Body       map[string]interface{}
}

type HttpClient interface {
    Post(endpoint string, payload map[string]interface{}) (*Response, error)
} 