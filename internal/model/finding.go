package model

type Severity string

const (
	Critical Severity = "CRITICAL"
	High     Severity = "HIGH"
	Medium   Severity = "MEDIUM"
	Low      Severity = "LOW"
	Info     Severity = "INFO"
)

type Finding struct {
	ID          string
	Severity    Severity
	Category    string
	File        string
	Line        int
	Title       string
	Description string
	Suggestion  string
}

type Results struct {
	FilesAnalyzed int
	LinesAnalyzed int
	Findings      []Finding
}

func (r Results) HasBlockingFindings() bool {
	for _, finding := range r.Findings {
		if finding.Severity == Critical {
			return true
		}
	}

	return false
}