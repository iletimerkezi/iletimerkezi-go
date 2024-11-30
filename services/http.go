package services

import "github.com/iletimerkezi/iletimerkezi-go/responses"

type HttpClient interface {
    Post(endpoint string, payload interface{}) (*responses.Response, error)
} 