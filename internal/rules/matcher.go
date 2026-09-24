package rules

import (
	"strings"
)

func Matches(line string, patterns []string) bool {
	lowerLine := strings.ToLower(line)

	for _, pattern := range patterns {
		if strings.Contains(
			lowerLine,
			strings.ToLower(pattern),
		) {
			return true
		}
	}

	return false
}