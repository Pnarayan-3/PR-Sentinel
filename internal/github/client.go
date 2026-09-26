package github

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetPullRequestFiles(
	pullNumber int,
) ([]PullRequestFile, error) {

	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/pulls/%d/files",
		c.Owner,
		c.Repository,
		pullNumber,
	)

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
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

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"GitHub API returned status %d",
			response.StatusCode,
		)
	}

	var files []PullRequestFile

	err = json.NewDecoder(
		response.Body,
	).Decode(&files)

	if err != nil {
		return nil, err
	}

	return files, nil
}