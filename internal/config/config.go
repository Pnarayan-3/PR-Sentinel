package config

type Config struct {
	RepositoryPath string
	RulesFile      string
	OutputFile     string

	// GitHub-related configuration.
	GitHubToken string
	Repository string
	PRNumber    string
}

type Environment struct {
	GitHubActions bool
}