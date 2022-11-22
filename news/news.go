package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient *http.Client
	key        string
	pageSize   int8
}

type Results struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	} `json:"articles"`
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

func (c *Client) FetchEverything(query, page string) (*Results, error) {
	encodedQuery := url.QueryEscape(query)
	endpoint := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%s&apiKey=%s&sortBy=publishedAt&language=en", encodedQuery, c.pageSize, page, c.key)
	resp, err := c.httpClient.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	res := &Results{}
	return res, json.Unmarshal(body, res)
}
