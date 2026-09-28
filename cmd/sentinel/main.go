package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Pnarayan-3/pr-sentinel/internal/analyzer"
	"github.com/Pnarayan-3/pr-sentinel/internal/config"
	githubclient "github.com/Pnarayan-3/pr-sentinel/internal/github"
	"github.com/Pnarayan-3/pr-sentinel/internal/report"
)

func main() {
	fmt.Println("🤖 PR Sentinel")
	fmt.Println("==========================")

	// Check whether PR Sentinel is running inside GitHub Actions
	// for a pull request.
	githubRepository := os.Getenv("GITHUB_REPOSITORY")
	prNumberString := os.Getenv("PR_NUMBER")
	// githubToken := os.Getenv("GITHUB_TOKEN")

	if githubRepository != "" &&
		prNumberString != "" &&
		githubToken != "" {

		runPullRequestMode(
			githubRepository,
			prNumberString,
			githubToken,
		)

		return
	}

	// Default behavior: analyze the local repository.
	runLocalMode()
}

func runPullRequestMode(
	repository string,
	prNumberString string,
	token string,
) {
	fmt.Println("Mode: GitHub Pull Request")
	fmt.Println()

	prNumber, err := strconv.Atoi(prNumberString)

	if err != nil {
		fmt.Printf(
			"Invalid PR_NUMBER: %v\n",
			err,
		)
		os.Exit(1)
	}

	parts := strings.SplitN(
		repository,
		"/",
		2,
	)

	if len(parts) != 2 {
		fmt.Printf(
			"Invalid GITHUB_REPOSITORY: %s\n",
			repository,
		)
		os.Exit(1)
	}

	owner := parts[0]
	repo := parts[1]

	fmt.Printf(
		"Repository: %s/%s\n",
		owner,
		repo,
	)

	fmt.Printf(
		"Pull Request: #%d\n",
		prNumber,
	)

	client := githubclient.NewClient(
		token,
		owner,
		repo,
	)

	files, err := client.GetChangedFiles(
		prNumberString,
	)

	if err != nil {
		fmt.Printf(
			"Failed to retrieve PR files: %v\n",
			err,
		)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf(
		"Changed files: %d\n",
		len(files),
	)

	for _, file := range files {
		fmt.Println()

		fmt.Printf(
			"📁 %s\n",
			file.Filename,
		)

		fmt.Printf(
			"   Status: %s\n",
			file.Status,
		)

		fmt.Printf(
			"   Added: %d\n",
			file.Additions,
		)

		fmt.Printf(
			"   Deleted: %d\n",
			file.Deletions,
		)

		fmt.Printf(
			"   Changes: %d\n",
			file.Changes,
		)

		if file.Patch != "" {
			fmt.Println(
				"   Patch available: yes",
			)
		} else {
			fmt.Println(
				"   Patch available: no",
			)
		}
	}

	fmt.Println()
	fmt.Println(
		"PR diff retrieval completed.",
	)
}

func runLocalMode() {
	fmt.Println("Mode: Local Repository")
	fmt.Println()

	cfg, err := config.Load(".")

	if err != nil {
		fmt.Printf(
			"Configuration error: %v\n",
			err,
		)
		os.Exit(1)
	}

	fmt.Printf(
		"Repository: %s\n",
		cfg.RepositoryPath,
	)

	fmt.Printf(
		"Rules file: %s\n",
		cfg.RulesFile,
	)

	results, err := analyzer.Analyze(cfg)

	if err != nil {
		fmt.Printf(
			"Analysis failed: %v\n",
			err,
		)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(
		"Analysis completed.",
	)

	fmt.Printf(
		"Files analyzed: %d\n",
		results.FilesAnalyzed,
	)

	fmt.Printf(
		"Findings: %d\n",
		len(results.Findings),
	)

	output := report.GenerateMarkdown(
		results,
	)

	fmt.Println()
	fmt.Println(output)

	if cfg.OutputFile != "" {
		if err := os.WriteFile(
			cfg.OutputFile,
			[]byte(output),
			0644,
		); err != nil {
			fmt.Printf(
				"Unable to write report: %v\n",
				err,
			)
			os.Exit(1)
		}

		fmt.Printf(
			"\nReport written to: %s\n",
			cfg.OutputFile,
		)
	}

	if results.HasBlockingFindings() {
		os.Exit(2)
	}
}