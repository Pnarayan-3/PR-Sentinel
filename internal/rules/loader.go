package rules

import (
	"os"
	"strings"
)

func LoadFromMarkdown(path string) ([]Rule, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(
		strings.ReplaceAll(
			string(data),
			"\r\n",
			"\n",
		),
		"\n",
	)

	var rules []Rule
	var current *Rule

	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "## ") {
			continue
		}

		if strings.HasPrefix(line, "### ") {
			if current != nil {
				rules = append(rules, *current)
			}

			title := strings.TrimSpace(
				strings.TrimPrefix(line, "### "),
			)

			current = &Rule{
				Name: title,
			}

			continue
		}

		if current == nil {
			continue
		}

		if strings.HasPrefix(line, "ID:") {
			current.ID = value(line)
		}

		if strings.HasPrefix(line, "Category:") {
			current.Category = value(line)
		}

		if strings.HasPrefix(line, "Severity:") {
			current.Severity = Severity(
				strings.ToUpper(value(line)),
			)
		}

		if strings.HasPrefix(line, "Description:") {
			current.Description = value(line)
		}

		if strings.HasPrefix(line, "Suggestion:") {
			current.Suggestion = value(line)
		}

		if strings.HasPrefix(line, "Pattern:") {
			pattern := value(line)

			if pattern != "" {
				current.Patterns = append(
					current.Patterns,
					pattern,
				)
			}
		}
	}

	if current != nil {
		rules = append(rules, *current)
	}

	return rules, nil
}

func value(line string) string {
	parts := strings.SplitN(line, ":", 2)

	if len(parts) != 2 {
		return ""
	}

	return strings.TrimSpace(parts[1])
}