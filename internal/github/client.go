package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Token      string
	Owner      string
	Repository string
}

func NewClient(
	token string,
	owner string,
	repository string,
) *Client {
	return &Client{
		Token:      token,
		Owner:      owner,
		Repository: repository,
	}
}

func (c *Client) request(
	method string,
	url string,
	body io.Reader,
) (*http.Response, error) {

	req, err := http.NewRequest(
		method,
		url,
		body,
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

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) ensureSuccess(
	response *http.Response,
) error {

	if response.StatusCode >= 200 &&
		response.StatusCode < 300 {
		return nil
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return fmt.Errorf(
			"GitHub API returned status %d",
			response.StatusCode,
		)
	}

	return fmt.Errorf(
		"GitHub API returned status %d: %s",
		response.StatusCode,
		string(body),
	)
}

func decodeJSON(
	response *http.Response,
	target interface{},
) error {

	defer response.Body.Close()

	return json.NewDecoder(
		response.Body,
	).Decode(target)
}