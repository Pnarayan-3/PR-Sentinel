package github

import (
	"fmt"
	"net/http"
)

type Comment struct {
	Body string `json:"body"`
}

func (c *Client) CreateIssueComment(
	number string,
	body string,
) error {
	path := fmt.Sprintf(
		"/repos/%s/issues/%s/comments",
		c.Repository,
		number,
	)

	response, err := c.request(
		http.MethodPost,
		path,
		Comment{
			Body: body,
		},
	)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	return c.ensureSuccess(response)
}