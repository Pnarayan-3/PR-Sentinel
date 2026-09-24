package report

import (
	"github.com/Pnarayan-3/pr-sentinel/internal/analyzer"
)

type Summary struct {
	Critical int
	High     int
	Medium   int
	Low      int
	Info     int
}

func BuildSummary(
	findings []analyzer.Finding,
) Summary {
	var summary Summary

	for _, finding := range findings {
		switch finding.Severity {
		case analyzer.Critical:
			summary.Critical++
		case analyzer.High:
			summary.High++
		case analyzer.Medium:
			summary.Medium++
		case analyzer.Low:
			summary.Low++
		case analyzer.Info:
			summary.Info++
		}
	}

	return summary
}