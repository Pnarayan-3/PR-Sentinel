package rules

import (
	"strings"

	"github.com/Pnarayan-3/pr-sentinel/internal/analyzer"
)

type Rule struct {
	ID          string
	Name        string
	Category    string
	Severity    Severity
	Description string
	Suggestion  string
	Patterns    []string
}

type Engine struct {
	Rules []Rule
}

func NewEngine(rules []Rule) *Engine {
	return &Engine{
		Rules: rules,
	}
}

func (e *Engine) AnalyzeFile(
	source analyzer.SourceFile,
) []analyzer.Finding {
	var findings []analyzer.Finding

	for lineNumber, line := range source.Lines {
		lowerLine := strings.ToLower(line)

		for _, rule := range e.Rules {
			for _, pattern := range rule.Patterns {
				if strings.Contains(
					lowerLine,
					strings.ToLower(pattern),
				) {
					findings = append(
						findings,
						analyzer.Finding{
							ID:          rule.ID,
							Severity:    analyzer.Severity(rule.Severity),
							Category:    rule.Category,
							File:        source.Path,
							Line:        lineNumber + 1,
							Title:       rule.Name,
							Description: rule.Description,
							Suggestion:  rule.Suggestion,
						},
					)

					break
				}
			}
		}
	}

	return findings
}