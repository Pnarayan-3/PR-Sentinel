package rules_test

import (
	"testing"

	"github.com/Pnarayan-3/pr-sentinel/internal/analyzer"
	"github.com/Pnarayan-3/pr-sentinel/internal/rules"
)

func TestEngineDetectsPattern(t *testing.T) {
	engine := rules.NewEngine([]rules.Rule{
		{
			ID:          "SEC-001",
			Name:        "Hardcoded Password",
			Category:    "Security",
			Severity:    rules.Critical,
			Description: "Password detected.",
			Suggestion:  "Use a secret manager.",
			Patterns:    []string{"password="},
		},
	})

	source := analyzer.SourceFile{
		Path: "config.go",
		Lines: []string{
			"package main",
			"password=hello123",
		},
	}

	findings := engine.AnalyzeFile(source)

	if len(findings) != 1 {
		t.Fatalf(
			"expected 1 finding, got %d",
			len(findings),
		)
	}

	if findings[0].ID != "SEC-001" {
		t.Fatalf(
			"expected SEC-001, got %s",
			findings[0].ID,
		)
	}

	if findings[0].Severity != analyzer.Critical {
		t.Fatalf(
			"expected CRITICAL, got %s",
			findings[0].Severity,
		)
	}
}