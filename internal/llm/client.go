package llm

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	APIKey  string
	BaseURL string
}

func NewClient(apiKey, baseURL string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

func (c *Client) MarshalRequest(request ChatRequest) ([]byte, error) {
	response, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) DoRequest(data []byte) (*http.Request, error) {
	req, err := http.NewRequest(
		http.MethodPost,
		c.BaseURL,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	return req, nil
}
