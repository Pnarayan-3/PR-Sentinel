package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	githubclient "github.com/Pnarayan-3/pr-sentinel/internal/github"
)

func main() {
	fmt.Println("🤖 PR Sentinel")
	fmt.Println("==========================")

	repository := os.Getenv("GITHUB_REPOSITORY")
	prNumberString := os.Getenv("PR_NUMBER")
	token := os.Getenv("GITHUB_TOKEN")

	if repository != "" &&
		prNumberString != "" &&
		token != "" {

		prNumber, err := strconv.Atoi(prNumberString)

		if err != nil {
			fmt.Println("Invalid PR_NUMBER")
			os.Exit(1)
		}

		parts := strings.SplitN(
			repository,
			"/",
			2,
		)

		if len(parts) != 2 {
			fmt.Println("Invalid GITHUB_REPOSITORY")
			os.Exit(1)
		}

		owner := parts[0]
		repo := parts[1]

		client := githubclient.NewClient(
			token,
			owner,
			repo,
		)

		files, err := client.GetPullRequestFiles(
			prNumber,
		)

		if err != nil {
			fmt.Printf(
				"Failed to retrieve PR files: %v\n",
				err,
			)

			os.Exit(1)
		}

		fmt.Printf(
			"PR #%d\n",
			prNumber,
		)

		fmt.Printf(
			"Changed files: %d\n",
			len(files),
		)

		for _, file := range files {
			fmt.Printf(
				"\n📁 %s\n",
				file.Filename,
			)

			fmt.Printf(
				"Status: %s\n",
				file.Status,
			)

			fmt.Printf(
				"Added: %d | Deleted: %d\n",
				file.Additions,
				file.Deletions,
			)
		}

		return
	}

	fmt.Println(
		"Running in local repository mode.",
	)

	// Existing local analyzer code goes here.
}