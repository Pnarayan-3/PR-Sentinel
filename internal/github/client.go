package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Token      string
	Repository string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(token, repository string) *Client {
	return &Client{
		Token:      token,
		Repository: repository,
		BaseURL:    "https://api.github.com",
		HTTPClient: &http.Client{},
	}
}

func (c *Client) request(
	method string,
	path string,
	body any,
) (*http.Response, error) {
	var requestBody *bytes.Reader

	if body != nil {
		data, err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

		requestBody = bytes.NewReader(data)
	} else {
		requestBody = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(
		method,
		c.BaseURL+path,
		requestBody,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+c.Token,
	)

	req.Header.Set(
		"Accept",
		"application/vnd.github+json",
	)

	req.Header.Set(
		"X-GitHub-Api-Version",
		"2022-11-28",
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	return c.HTTPClient.Do(req)
}

func (c *Client) ensureSuccess(response *http.Response) error {
	if response.StatusCode >= 200 &&
		response.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf(
		"github API returned status %d",
		response.StatusCode,
	)
}