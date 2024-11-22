package services

import (
	"fmt"
	"time"
)

type SummaryService struct {
	client        HttpClient
	apiKey        string
	apiHash       string
	lastStartDate time.Time
	lastEndDate   time.Time
	lastPage      int
}

func NewSummaryService(client HttpClient, apiKey, apiHash string) *SummaryService {
	return &SummaryService{
		client:  client,
		apiKey:  apiKey,
		apiHash: apiHash,
	}
}

func (s *SummaryService) List(startDate, endDate time.Time, page int) (*SummaryResponse, error) {
	s.lastStartDate = startDate
	s.lastEndDate = endDate
	s.lastPage = page

	payload := map[string]interface{}{
		"request": map[string]interface{}{
			"authentication": map[string]string{
				"key":  s.apiKey,
				"hash": s.apiHash,
			},
			"filter": map[string]interface{}{
				"start": startDate.Format("2006-01-02"),
				"end":   endDate.Format("2006-01-02"),
				"page":  page,
			},
		},
	}

	resp, err := s.client.Post("get-reports/json", payload)
	if err != nil {
		return nil, err
	}

	return NewSummaryResponse(resp), nil
}

func (s *SummaryService) Next() (*SummaryResponse, error) {
	if s.lastPage == 0 {
		return nil, fmt.Errorf("no previous summary request found. Call List() first")
	}

	return s.List(s.lastStartDate, s.lastEndDate, s.lastPage+1)
} 