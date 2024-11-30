package services

import (
	"fmt"
	"github.com/iletimerkezi/iletimerkezi-go/responses"
)

type ReportService struct {
	client      HttpClient
	apiKey      string
	apiHash     string
	lastOrderID int
	lastPage    int
}

func NewReportService(client HttpClient, apiKey, apiHash string) *ReportService {
	return &ReportService{
		client:  client,
		apiKey:  apiKey,
		apiHash: apiHash,
	}
}

func (s *ReportService) Get(orderID, page int, rowCount int) (*responses.ReportResponse, error) {
	s.lastOrderID = orderID
	s.lastPage = page

	payload := map[string]interface{}{
		"request": map[string]interface{}{
			"authentication": map[string]string{
				"key":  s.apiKey,
				"hash": s.apiHash,
			},
			"order": map[string]interface{}{
				"id":       orderID,
				"page":     page,
				"rowCount": rowCount,
			},
		},
	}

	resp, err := s.client.Post("get-report/json", payload)
	if err != nil {
		return nil, err
	}

	return responses.NewReportResponse(resp), nil
}

func (s *ReportService) Next() (*responses.ReportResponse, error) {
	if s.lastPage == 0 {
		return nil, fmt.Errorf("no previous report request found. Call Get() first")
	}

	return s.Get(s.lastOrderID, s.lastPage+1, 1000)
} 