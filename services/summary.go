package services

import (
    "fmt"
    "time"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
)

type SummaryService struct {
    client        HttpClient
    apiKey        string
    apiHash       string
    lastStartDate *time.Time
    lastEndDate   *time.Time
    lastPage      int
}

func NewSummaryService(client HttpClient, apiKey, apiHash string) *SummaryService {
    return &SummaryService{
        client:  client,
        apiKey:  apiKey,
        apiHash: apiHash,
    }
}

func (s *SummaryService) Get(startDate, endDate *time.Time, page int) (*responses.SummaryResponse, error) {
    if page < 1 {
        page = 1
    }

    // Tarihleri sakla
    s.lastStartDate = startDate
    s.lastEndDate = endDate
    s.lastPage = page

    return s.List(startDate, endDate, page)
}

func (s *SummaryService) List(startDate, endDate *time.Time, page int) (*responses.SummaryResponse, error) {
    payload := map[string]interface{}{
        "request": map[string]interface{}{
            "authentication": map[string]string{
                "key":  s.apiKey,
                "hash": s.apiHash,
            },
            "filter": map[string]interface{}{
                "page": page,
            },
        },
    }

    // Tarih filtrelerini ekle
    if startDate != nil || endDate != nil {
        filter := payload["request"].(map[string]interface{})["filter"].(map[string]interface{})
        if startDate != nil {
            filter["start"] = startDate.Format("2006-01-02")
        }
        if endDate != nil {
            filter["end"] = endDate.Format("2006-01-02")
        }
    }

    resp, err := s.client.Post("get-report-summary/json", payload)
    if err != nil {
        return nil, err
    }

    return responses.NewSummaryResponse(resp), nil
}

func (s *SummaryService) Next() (*responses.SummaryResponse, error) {
    if s.lastPage == 0 {
        return nil, fmt.Errorf("no previous summary request found. Call Get() first")
    }

    return s.List(s.lastStartDate, s.lastEndDate, s.lastPage+1)
}