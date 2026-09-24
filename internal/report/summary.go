package report

import "github.com/Pnarayan-3/pr-sentinel/internal/model"

type Summary struct {
	Critical int
	High     int
	Medium   int
	Low      int
	Info     int
}

func BuildSummary(
	findings []model.Finding,
) Summary {
	var summary Summary

	for _, finding := range findings {
		switch finding.Severity {
		case model.Critical:
			summary.Critical++
		case model.High:
			summary.High++
		case model.Medium:
			summary.Medium++
		case model.Low:
			summary.Low++
		case model.Info:
			summary.Info++
		}
	}

	return summary
}