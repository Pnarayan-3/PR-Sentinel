package diff

import (
	"strconv"
	"strings"
)

type ChangedLine struct {
	LineNumber int
	Content    string
}

type FileDiff struct {
	Filename     string
	ChangedLines []ChangedLine
}

func ParsePatch(
	filename string,
	patch string,
) FileDiff {

	result := FileDiff{
		Filename: filename,
	}

	lines := strings.Split(
		patch,
		"\n",
	)

	currentLine := 0

	for _, line := range lines {

		if strings.HasPrefix(
			line,
			"@@",
		) {
			currentLine = parseNewLineNumber(line)
			continue
		}

		if strings.HasPrefix(
			line,
			"+++",
		) {
			continue
		}

		if strings.HasPrefix(
			line,
			"+",
		) {
			result.ChangedLines = append(
				result.ChangedLines,
				ChangedLine{
					LineNumber: currentLine,
					Content:    strings.TrimPrefix(
						line,
						"+",
					),
				},
			)

			currentLine++
			continue
		}

		if strings.HasPrefix(
			line,
			"-",
		) {
			continue
		}

		if strings.HasPrefix(
			line,
			"\\ No newline",
		) {
			continue
		}

		currentLine++
	}

	return result
}

func parseNewLineNumber(
	header string,
) int {

	parts := strings.Fields(
		header,
	)

	for _, part := range parts {

		if !strings.HasPrefix(
			part,
			"+",
		) {
			continue
		}

		part = strings.TrimPrefix(
			part,
			"+",
		)

		part = strings.TrimSuffix(
			part,
			"@@",
		)

		if strings.Contains(
			part,
			",",
		) {
			part = strings.Split(
				part,
				",",
			)[0]
		}

		number, err := strconv.Atoi(
			part,
		)

		if err == nil {
			return number
		}
	}

	return 0
}