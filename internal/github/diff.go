package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ChangedFile struct {
	Filename  string `json:"filename"`
	Status    string `json:"status"`
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
	Changes   int    `json:"changes"`
	Patch     string `json:"patch"`
}

func (c *Client) GetChangedFiles(
	number string,
) ([]ChangedFile, error) {
	path := fmt.Sprintf(
		"/repos/%s/%s/pulls/%s/files",
		c.Owner,
		c.Repository,
		number,
	)

	response, err := c.request(
		http.MethodGet,
		path,
		nil,
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if err := c.ensureSuccess(response); err != nil {
		return nil, err
	}

	var result []ChangedFile

	if err := json.NewDecoder(
		response.Body,
	).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}