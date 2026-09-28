package config

import (
	"os"
	"path/filepath"
)

func Load(repositoryPath string) (Config, error) {
	absolutePath, err := filepath.Abs(repositoryPath)
	if err != nil {
		return Config{}, err
	}

	rulesFile := filepath.Join(
		absolutePath,
		"PR-SENTINEL.md",
	)

	outputFile := filepath.Join(
		absolutePath,
		"pr-sentinel-report.md",
	)

	return Config{
		RepositoryPath: absolutePath,
		RulesFile:      rulesFile,
		OutputFile:     outputFile,

		// GitHubToken: os.Getenv("GITHUB_TOKEN"),
		Repository:  os.Getenv("GITHUB_REPOSITORY"),
		PRNumber:    os.Getenv("PR_NUMBER"),
	}, nil
}

func LoadEnvironment() Environment {
	return Environment{
		GitHubActions: os.Getenv("GITHUB_ACTIONS") == "true",
	}
}