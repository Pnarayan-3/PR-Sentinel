package analyzer

import (
	"fmt"
	"path/filepath"

	"github.com/Pnarayan-3/pr-sentinel/internal/config"
	"github.com/Pnarayan-3/pr-sentinel/internal/model"
	"github.com/Pnarayan-3/pr-sentinel/internal/rules"
)

func Analyze(cfg config.Config) (model.Results, error) {
	ruleDefinitions, err := rules.LoadFromMarkdown(
		cfg.RulesFile,
	)

	if err != nil {
		return model.Results{}, fmt.Errorf(
			"failed to load rules: %w",
			err,
		)
	}

	engine := rules.NewEngine(ruleDefinitions)

	files, err := ScanRepository(
		cfg.RepositoryPath,
	)

	if err != nil {
		return model.Results{}, fmt.Errorf(
			"failed to scan repository: %w",
			err,
		)
	}

	results := model.Results{}

	for _, file := range files {
		source, err := ReadSourceFile(file)
		if err != nil {
			continue
		}

		results.FilesAnalyzed++
		results.LinesAnalyzed += len(source.Lines)

		findings := engine.AnalyzeFile(source)

		for i := range findings {
			findings[i].File = relativePath(
				cfg.RepositoryPath,
				findings[i].File,
			)
		}

		results.Findings = append(
			results.Findings,
			findings...,
		)
	}

	return results, nil
}

func relativePath(root, file string) string {
	relative, err := filepath.Rel(root, file)

	if err != nil {
		return file
	}

	return filepath.ToSlash(relative)
}