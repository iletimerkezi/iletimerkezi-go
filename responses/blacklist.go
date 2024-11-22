package responses

type BlacklistResponse struct {
    *BaseResponse
    numbers  []string
    position int
}

func NewBlacklistResponse(resp *Response) *BlacklistResponse {
    base := NewBaseResponse(resp)
    r := &BlacklistResponse{
        BaseResponse: base,
    }
    r.parseNumbers()
    return r
}

func (r *BlacklistResponse) parseNumbers() {
    if blacklist, ok := r.Data["blacklist"].(map[string]interface{}); ok {
        if numbers, ok := blacklist["number"].([]interface{}); ok {
            r.numbers = make([]string, len(numbers))
            for i, num := range numbers {
                if str, ok := num.(string); ok {
                    r.numbers[i] = str
                }
            }
        }
    }
}

func (r *BlacklistResponse) GetCount() int {
    if blacklist, ok := r.Data["blacklist"].(map[string]interface{}); ok {
        if count, ok := blacklist["count"].(float64); ok {
            return int(count)
        }
    }
    return 0
}

func (r *BlacklistResponse) GetNumbers() []string {
    return r.numbers
}

// Iterator implementation
func (r *BlacklistResponse) Next() bool {
    r.position++
    return r.position < len(r.numbers)
}

func (r *BlacklistResponse) Current() string {
    if r.position >= len(r.numbers) {
        return ""
    }
    return r.numbers[r.position]
}

func (r *BlacklistResponse) Reset() {
    r.position = 0
} 