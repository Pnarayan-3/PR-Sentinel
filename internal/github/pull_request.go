package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Draft  bool   `json:"draft"`
	State  string `json:"state"`
}

func (c *Client) GetPullRequest(
	number string,
) (PullRequest, error) {
	path := fmt.Sprintf(
		"/repos/%s/pulls/%s",
		c.Repository,
		number,
	)

	response, err := c.request(
		http.MethodGet,
		path,
		nil,
	)

	if err != nil {
		return PullRequest{}, err
	}

	defer response.Body.Close()

	if err := c.ensureSuccess(response); err != nil {
		return PullRequest{}, err
	}

	var pullRequest PullRequest

	if err := json.NewDecoder(
		response.Body,
	).Decode(&pullRequest); err != nil {
		return PullRequest{}, err
	}

	return pullRequest, nil
}