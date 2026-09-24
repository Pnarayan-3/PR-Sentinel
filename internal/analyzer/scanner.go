package analyzer

import (
	"os"
	"path/filepath"
	"strings"
)

var ignoredDirectories = map[string]bool{
	".git":         true,
	".github":      true,
	".idea":        true,
	".vscode":      true,
	"node_modules": true,
	"vendor":       true,
	"target":       true,
	"build":        true,
	"dist":         true,
	".terraform":   true,
}

func ScanRepository(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(
		path string,
		info os.FileInfo,
		err error,
	) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if ignoredDirectories[info.Name()] {
				return filepath.SkipDir
			}

			return nil
		}

		if !isAnalyzableFile(path) {
			return nil
		}

		files = append(files, path)

		return nil
	})

	return files, err
}

func isAnalyzableFile(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))

	switch extension {
	case ".go",
		".java",
		".cs",
		".py",
		".js",
		".ts",
		".tsx",
		".jsx",
		".rb",
		".php",
		".c",
		".cpp",
		".h",
		".hpp",
		".yaml",
		".yml",
		".json",
		".tf",
		".sh",
		".sql":
		return true
	}

	return false
}