package llm

import (
	"encoding/json"
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