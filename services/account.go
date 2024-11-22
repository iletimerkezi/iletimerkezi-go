package services

type AccountService struct {
    client  HttpClient
    apiKey  string
    apiHash string
}

func NewAccountService(client HttpClient, apiKey, apiHash string) *AccountService {
    return &AccountService{
        client:  client,
        apiKey:  apiKey,
        apiHash: apiHash,
    }
}

func (s *AccountService) Balance() (*AccountResponse, error) {
    payload := map[string]interface{}{
        "request": map[string]interface{}{
            "authentication": map[string]string{
                "key":  s.apiKey,
                "hash": s.apiHash,
            },
        },
    }

    resp, err := s.client.Post("get-balance/json", payload)
    if err != nil {
        return nil, err
    }

    return NewAccountResponse(resp), nil
} 