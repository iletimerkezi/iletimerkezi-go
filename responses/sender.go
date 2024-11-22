package responses

type SenderResponse struct {
    *BaseResponse
    senders  []string
    position int
}

func NewSenderResponse(resp *Response) *SenderResponse {
    base := NewBaseResponse(resp)
    r := &SenderResponse{
        BaseResponse: base,
    }
    r.parseSenders()
    return r
}

func (r *SenderResponse) parseSenders() {
    if senders, ok := r.Data["senders"].(map[string]interface{}); ok {
        if list, ok := senders["sender"].([]interface{}); ok {
            r.senders = make([]string, len(list))
            for i, sender := range list {
                if str, ok := sender.(string); ok {
                    r.senders[i] = str
                }
            }
        }
    }
}

func (r *SenderResponse) GetSenders() []string {
    return r.senders
}

// Iterator implementation
func (r *SenderResponse) Next() bool {
    r.position++
    return r.position < len(r.senders)
}

func (r *SenderResponse) Current() string {
    if r.position >= len(r.senders) {
        return ""
    }
    return r.senders[r.position]
}

func (r *SenderResponse) Reset() {
    r.position = 0
} 