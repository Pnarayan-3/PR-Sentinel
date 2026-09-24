package analyzer

import (
	"os"
	"strings"

	"github.com/Pnarayan-3/pr-sentinel/internal/model"
)

func ReadSourceFile(path string) (model.SourceFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.SourceFile{}, err
	}

	content := strings.ReplaceAll(
		string(data),
		"\r\n",
		"\n",
	)

	lines := strings.Split(content, "\n")

	return model.SourceFile{
		Path:  path,
		Lines: lines,
	}, nil
}