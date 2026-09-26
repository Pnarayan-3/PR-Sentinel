package github

type PullRequestFile struct {
	Filename string `json:"filename"`
	Status   string `json:"status"`
	Patch    string `json:"patch"`
	Additions int   `json:"additions"`
	Deletions int   `json:"deletions"`
	Changes   int   `json:"changes"`
}