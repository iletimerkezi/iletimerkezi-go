package services

import (
	"time"
)

type BlacklistService struct {
	client  HttpClient
	apiKey  string
	apiHash string
}

func NewBlacklistService(client HttpClient, apiKey, apiHash string) *BlacklistService {
	return &BlacklistService{
		client:  client,
		apiKey:  apiKey,
		apiHash: apiHash,
	}
}

func (s *BlacklistService) List(startDate, endDate *time.Time, page int, rowCount int) (*BlacklistResponse, error) {
	payload := map[string]interface{}{
		"request": map[string]interface{}{
			"authentication": map[string]string{
				"key":  s.apiKey,
				"hash": s.apiHash,
			},
			"blacklist": map[string]interface{}{
				"page":     page,
				"rowCount": min(rowCount, 1000),
			},
		},
	}

	if startDate != nil || endDate != nil {
		filter := make(map[string]string)
		if startDate != nil {
			filter["start"] = startDate.Format("2006-01-02 15:04:05")
		}
		if endDate != nil {
			filter["end"] = endDate.Format("2006-01-02 15:04:05")
		}
		payload["request"].(map[string]interface{})["blacklist"].(map[string]interface{})["filter"] = filter
	}

	resp, err := s.client.Post("get-blacklist/json", payload)
	if err != nil {
		return nil, err
	}

	return NewBlacklistResponse(resp), nil
} 