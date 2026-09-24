package analyzer

import (
	"os"
	"path/filepath"
	"strings"
)

func ScanRepository(root string) ([]string, error) {
	var files []string

	err := filepath.Walk(
		root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Ignore directories that should never be analyzed.
			if info.IsDir() {
				if shouldIgnoreDirectory(path, root) {
					return filepath.SkipDir
				}

				return nil
			}

			// Ignore files that should not be analyzed.
			if shouldIgnoreFile(path, root) {
				return nil
			}

			files = append(files, path)

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func shouldIgnoreDirectory(path, root string) bool {
	relativePath, err := filepath.Rel(root, path)
	if err != nil {
		return false
	}

	relativePath = filepath.ToSlash(relativePath)

	if relativePath == "." {
		return false
	}

	ignoredDirectories := []string{
		".git",
		".github",
		"node_modules",
		"vendor",
		"testdata",
	}

	for _, directory := range ignoredDirectories {
		if relativePath == directory ||
			strings.HasPrefix(relativePath, directory+"/") {

			return true
		}
	}

	return false
}

func shouldIgnoreFile(path, root string) bool {
	relativePath, err := filepath.Rel(root, path)
	if err != nil {
		return false
	}

	relativePath = filepath.ToSlash(relativePath)

	// Ignore files inside testdata directories.
	if strings.Contains(relativePath, "/testdata/") ||
		strings.HasPrefix(relativePath, "testdata/") {
		return true
	}

	return false
}