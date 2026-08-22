package llm

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"errors"
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

func (c *Client) SendRequest(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ReadResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) ParseResponse(body []byte) (*ChatResponse, error) {
	var response ChatResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) ExtractContent(response *ChatResponse) (string, error) {
	if len(response.Choices) == 0 {
		return "", errors.New("no choices in response")
	}

	return response.Choices[0].Message.Content, nil
}