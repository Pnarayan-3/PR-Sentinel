package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const githubAPIBaseURL = "https://api.github.com"

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
	path string,
	body interface{},
) (*http.Response, error) {

	var requestBody io.Reader

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to encode request body: %w",
				err,
			)
		}

		requestBody = bytes.NewBuffer(jsonBody)
	}

	url := githubAPIBaseURL + path

	req, err := http.NewRequest(
		method,
		url,
		requestBody,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"failed to create GitHub request: %w",
			err,
		)
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

	if body != nil {
		req.Header.Set(
			"Content-Type",
			"application/json",
		)
	}

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf(
			"GitHub API request failed: %w",
			err,
		)
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