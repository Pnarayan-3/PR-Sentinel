package main

import (
	"fmt"
	"os"

	"github.com/Pnarayan-3/pr-sentinel/internal/analyzer"
	"github.com/Pnarayan-3/pr-sentinel/internal/config"
	"github.com/Pnarayan-3/pr-sentinel/internal/report"
)

func main() {
	fmt.Println("🤖 PR Sentinel")
	fmt.Println("==========================")

	cfg, err := config.Load(".")
	if err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Repository: %s\n", cfg.RepositoryPath)
	fmt.Printf("Rules file: %s\n", cfg.RulesFile)

	results, err := analyzer.Analyze(cfg)
	if err != nil {
		fmt.Printf("Analysis failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Analysis completed.")
	fmt.Printf("Files analyzed: %d\n", results.FilesAnalyzed)
	fmt.Printf("Findings: %d\n", len(results.Findings))

	output := report.GenerateMarkdown(results)

	fmt.Println()
	fmt.Println(output)

	if cfg.OutputFile != "" {
		if err := os.WriteFile(
			cfg.OutputFile,
			[]byte(output),
			0644,
		); err != nil {
			fmt.Printf("Unable to write report: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nReport written to: %s\n", cfg.OutputFile)
	}

	if results.HasBlockingFindings() {
		os.Exit(2)
	}
}