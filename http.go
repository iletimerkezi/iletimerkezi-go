package iletimerkezi

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "github.com/iletimerkezi/iletimerkezi-go/responses"
)

type HttpClient interface {
    Post(url string, payload interface{}) (*responses.Response, error)
    GetLastResponse() *responses.Response
    GetLastPayload() []byte
}

type defaultHttpClient struct {
    client       *http.Client
    lastResponse *responses.Response
    lastPayload  []byte
    baseURL      string
}

func NewHttpClient() HttpClient {
    return &defaultHttpClient{
        client:  &http.Client{},
        baseURL: "https://api.iletimerkezi.com/v1/",
    }
}

func (c *defaultHttpClient) Post(url string, payload interface{}) (*responses.Response, error) {
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal payload: %w", err)
    }
    c.lastPayload = jsonData

    req, err := http.NewRequest("POST", c.baseURL+url, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    var responseBody map[string]interface{}
    if err := json.Unmarshal(body, &responseBody); err != nil {
        return nil, fmt.Errorf("failed to parse response body: %w", err)
    }

    response := &responses.Response{
        StatusCode: resp.StatusCode,
        Body:      responseBody,
    }

    c.lastResponse = response
    return response, nil
}

func (c *defaultHttpClient) GetLastResponse() *responses.Response {
    return c.lastResponse
}

func (c *defaultHttpClient) GetLastPayload() []byte {
    return c.lastPayload
} 