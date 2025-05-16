package httpclient

import (
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{httpClient: &http.Client{}}
}

func (c *Client) DoRequest(url string) (int, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
