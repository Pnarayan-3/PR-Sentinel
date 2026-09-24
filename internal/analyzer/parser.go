package analyzer

import (
	"os"
	"strings"
)

type SourceFile struct {
	Path  string
	Lines []string
}

func ReadSourceFile(path string) (SourceFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return SourceFile{}, err
	}

	content := strings.ReplaceAll(
		string(data),
		"\r\n",
		"\n",
	)

	lines := strings.Split(content, "\n")

	return SourceFile{
		Path:  path,
		Lines: lines,
	}, nil
}