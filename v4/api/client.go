package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

// New returns an authenticated API Client
func New(options ...Option) (*Client, error) {
	config := &Config{}
	for _, option := range options {
		option(config)
	}

	httpClient, err := HTTPClient(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		http:   httpClient,
		config: config,
	}, nil
}

// Client is the owner of all communication with IMGIX
type Client struct {
	http   *http.Client
	config *Config
}

func (c *Client) do(req *http.Request, into interface{}) error {
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return errors.New(resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(into)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
