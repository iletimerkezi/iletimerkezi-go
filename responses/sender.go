package responses

type SenderResponse struct {
    *BaseResponse
    Senders  []string
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
            r.Senders = make([]string, len(list))
            for i, sender := range list {
                if str, ok := sender.(string); ok {
                    r.Senders[i] = str
                }
            }
        }
    }
}