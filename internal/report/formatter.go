package report

import (
	"fmt"
	"strings"

	"github.com/Pnarayan-3/pr-sentinel/internal/analyzer"
)

func severityIcon(severity analyzer.Severity) string {
	switch severity {
	case analyzer.Critical:
		return "🔴"
	case analyzer.High:
		return "🟠"
	case analyzer.Medium:
		return "🟡"
	case analyzer.Low:
		return "🔵"
	default:
		return "ℹ️"
	}
}

func formatFinding(finding analyzer.Finding) string {
	var builder strings.Builder

	builder.WriteString(
		fmt.Sprintf(
			"### %s %s — %s\n\n",
			severityIcon(finding.Severity),
			finding.Severity,
			finding.ID,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"**%s**\n\n",
			finding.Title,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"📁 `%s`  \n📍 Line `%d`\n\n",
			finding.File,
			finding.Line,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"**Description:** %s\n\n",
			finding.Description,
		),
	)

	builder.WriteString(
		fmt.Sprintf(
			"**Suggestion:** %s\n\n",
			finding.Suggestion,
		),
	)

	return builder.String()
}