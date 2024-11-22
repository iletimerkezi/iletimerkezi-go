package iletimerkezi

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type HttpClient interface {
    Post(url string, payload interface{}) (*Response, error)
    GetLastResponse() *Response
}

type defaultHttpClient struct {
    client        *http.Client
    lastResponse  *Response
    lastPayload   []byte
    baseURL       string
    userAgent     string
}

type Response struct {
    StatusCode int
    Body       map[string]interface{}
    RawBody    []byte
}

func NewDefaultHttpClient() HttpClient {
    return &defaultHttpClient{
        client:    &http.Client{},
        baseURL:   "https://api.iletimerkezi.com/v1/",
        userAgent: fmt.Sprintf("IletiMerkezi-Go/%s", Version()),
    }
}

func (c *defaultHttpClient) Post(url string, payload interface{}) (*Response, error) {
    jsonData, err := json.Marshal(payload)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal payload: %w", err)
    }

    c.lastPayload = jsonData
    fullURL := c.baseURL + url

    req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", c.userAgent)

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    var result map[string]interface{}
    if err := json.Unmarshal(body, &result); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }

    response := &Response{
        StatusCode: resp.StatusCode,
        Body:       result,
        RawBody:    body,
    }

    c.lastResponse = response
    return response, nil
}

func (c *defaultHttpClient) GetLastResponse() *Response {
    return c.lastResponse
} 