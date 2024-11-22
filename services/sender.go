package services

type SenderService struct {
    client  HttpClient
    apiKey  string
    apiHash string
}

func NewSenderService(client HttpClient, apiKey, apiHash string) *SenderService {
    return &SenderService{
        client:  client,
        apiKey:  apiKey,
        apiHash: apiHash,
    }
}

func (s *SenderService) List() (*SenderResponse, error) {
    payload := map[string]interface{}{
        "request": map[string]interface{}{
            "authentication": map[string]string{
                "key":  s.apiKey,
                "hash": s.apiHash,
            },
        },
    }

    resp, err := s.client.Post("get-sender/json", payload)
    if err != nil {
        return nil, err
    }

    return NewSenderResponse(resp), nil
} 