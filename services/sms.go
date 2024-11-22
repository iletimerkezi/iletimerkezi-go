package services

import (
    "time"
)

type SmsService struct {
    client      HttpClient
    apiKey      string
    apiHash     string
    sender      string
    sendTime    string
    iysEnabled  bool
    iysList     string
}

func NewSmsService(client HttpClient, apiKey, apiHash, sender string) *SmsService {
    return &SmsService{
        client:     client,
        apiKey:     apiKey,
        apiHash:    apiHash,
        sender:     sender,
        iysEnabled: true,
        iysList:    "BIREYSEL",
    }
}

func (s *SmsService) Schedule(t time.Time) *SmsService {
    s.sendTime = t.Format("2006-01-02 15:04:05")
    return s
}

func (s *SmsService) EnableIysConsent() *SmsService {
    s.iysEnabled = true
    return s
}

func (s *SmsService) DisableIysConsent() *SmsService {
    s.iysEnabled = false
    return s
}

func (s *SmsService) SetIysList(list string) *SmsService {
    s.iysList = list
    return s
}

func (s *SmsService) Send(recipients interface{}, message string) (*SmsResponse, error) {
    payload := map[string]interface{}{
        "request": map[string]interface{}{
            "authentication": map[string]string{
                "key":  s.apiKey,
                "hash": s.apiHash,
            },
            "order": map[string]interface{}{
                "sender":       s.sender,
                "sendDateTime": s.sendTime,
                "iys":         s.iysEnabled,
                "iysList":     s.iysList,
                "message":     s.buildMessages(recipients, message),
            },
        },
    }

    resp, err := s.client.Post("send-sms/json", payload)
    if err != nil {
        return nil, err
    }

    return NewSmsResponse(resp), nil
}

func (s *SmsService) Cancel(orderID string) (*SmsResponse, error) {
    payload := map[string]interface{}{
        "request": map[string]interface{}{
            "authentication": map[string]string{
                "key":  s.apiKey,
                "hash": s.apiHash,
            },
            "order": map[string]string{
                "id": orderID,
            },
        },
    }

    resp, err := s.client.Post("cancel-order/json", payload)
    if err != nil {
        return nil, err
    }

    return NewSmsResponse(resp), nil
}

func (s *SmsService) buildMessages(recipients interface{}, message string) interface{} {
    switch v := recipients.(type) {
    case string:
        return map[string]interface{}{
            "text": message,
            "receipents": map[string][]string{
                "number": {v},
            },
        }
    case []string:
        return map[string]interface{}{
            "text": message,
            "receipents": map[string][]string{
                "number": v,
            },
        }
    case map[string]string:
        messages := make([]map[string]interface{}, 0)
        for number, text := range v {
            messages = append(messages, map[string]interface{}{
                "text": text,
                "receipents": map[string][]string{
                    "number": {number},
                },
            })
        }
        return messages
    default:
        return nil
    }
} 