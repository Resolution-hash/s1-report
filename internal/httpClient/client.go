package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}
func (c *Client) Post(endpoint string, body interface{}, headers map[string]string) (int, []byte, error) {
	url := c.BaseURL + endpoint

	jsonData, err := json.Marshal(body)
	if err != nil {
		return 0, nil, fmt.Errorf("error marshalling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, nil, fmt.Errorf("error sending request %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("error reading response: %w", err)
	}

	return resp.StatusCode, bodyBytes, nil
}

func (c *Client) Get(endpoint string, headers map[string]string) (int, []byte, error) {
	url := c.BaseURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, nil, fmt.Errorf("error sending request %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("error reading response: %w", err)
	}

	return resp.StatusCode, bodyBytes, nil
}
