package news

import "net/http"

type Client struct {
	httpClient *http.Client
	key        string
	pageSize   int8
}

func DefaultNewsClient(apiKey string) *Client {
	return NewNewsClient(apiKey, 100, http.DefaultClient)
}

func NewNewsClient(apiKey string, pageSize int8, httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
		key:        apiKey,
		pageSize:   pageSize,
	}
}
