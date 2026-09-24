package report

import (
	"fmt"
	"strings"

	"github.com/Pnarayan-3/pr-sentinel/internal/model"
)

func GenerateMarkdown(results model.Results) string {
	summary := BuildSummary(results.Findings)

	var builder strings.Builder

	builder.WriteString("# 🤖 PR Sentinel Review\n\n")

	builder.WriteString(
		"PR Sentinel is a rule-based code review agent.\n\n",
	)

	builder.WriteString("## Summary\n\n")

	builder.WriteString(
		fmt.Sprintf(
			"| Metric | Count |\n|---|---:|\n| Files analyzed | %d |\n| Lines analyzed | %d |\n| 🔴 Critical | %d |\n| 🟠 High | %d |\n| 🟡 Medium | %d |\n| 🔵 Low | %d |\n| ℹ️ Info | %d |\n\n",
			results.FilesAnalyzed,
			results.LinesAnalyzed,
			summary.Critical,
			summary.High,
			summary.Medium,
			summary.Low,
			summary.Info,
		),
	)

	if len(results.Findings) == 0 {
		builder.WriteString(
			"## ✅ No issues found\n\n",
		)

		builder.WriteString(
			"PR Sentinel did not detect any configured rule violations.\n",
		)

		return builder.String()
	}

	builder.WriteString("## Findings\n\n")

	for _, finding := range results.Findings {
		builder.WriteString(
			formatFinding(finding),
		)

		builder.WriteString("---\n\n")
	}

	if results.HasBlockingFindings() {
		builder.WriteString(
			"## ❌ Review Status\n\n",
		)

		builder.WriteString(
			"**FAILED** — at least one critical issue was detected.\n",
		)
	} else {
		builder.WriteString(
			"## ✅ Review Status\n\n",
		)

		builder.WriteString(
			"**PASSED** — no critical issues were detected.\n",
		)
	}

	return builder.String()
}