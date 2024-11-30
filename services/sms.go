package services

import (
    "time"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
)

const (
    IysListBireysel = "BIREYSEL"
    IysListTacir    = "TACIR"
)

type SmsService struct {
    client      HttpClient
    apiKey      string
    apiHash     string
    sender      string
    schedule    string
    iys         string
    iysList     string
}

func NewSmsService(client HttpClient, apiKey, apiHash, sender string) *SmsService {
    return &SmsService{
        client:     client,
        apiKey:     apiKey,
        apiHash:    apiHash,
        sender:     sender,
		schedule:   "",
        iys:        "1",
        iysList:    IysListBireysel,
    }
}

func (s *SmsService) Schedule(t time.Time) *SmsService {
    s.schedule = t.Format("02/01/2006 15:04")
    return s
}

func (s *SmsService) EnableIysConsent() *SmsService {
    s.iys = "1"
    return s
}

func (s *SmsService) DisableIysConsent() *SmsService {
    s.iys = "0"
    return s
}

func (s *SmsService) SetIysList(list string) *SmsService {
    if list != IysListBireysel && list != IysListTacir {
        return s
    }
    s.iysList = list
    return s
}

func (s *SmsService) Send(recipients interface{}, message string, sender string) (*responses.SmsResponse, error) {
    
	if sender != "" {
		s.sender = sender
	}

	payload := map[string]interface{}{
        "request": map[string]interface{}{
            "authentication": map[string]string{
                "key":  s.apiKey,
                "hash": s.apiHash,
            },
			"order": map[string]interface{}{
				"sender": s.sender,
				"sendDateTime": s.schedule,
				"iys": s.iys,
				"iysList": s.iysList,
				"message": s.buildMessages(recipients, message),
			},
        },
    }

    resp, err := s.client.Post("send-sms/json", payload)
    if err != nil {
        return nil, err
    }

    return responses.NewSmsResponse(resp), nil
}

func (s *SmsService) Cancel(orderID string) (*responses.SmsResponse, error) {
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

    return responses.NewSmsResponse(resp), nil
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