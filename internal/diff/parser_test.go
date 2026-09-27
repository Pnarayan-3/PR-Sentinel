package diff

import (
	"testing"
)

func TestParsePatch(t *testing.T) {

	patch := `@@ -10,4 +10,6 @@
 func login() {
     username := "admin"
+    password := "hello123"
+    debug := true
     return
 }`

	result := ParsePatch(
		"config.go",
		patch,
	)

	if result.Filename != "config.go" {
		t.Fatalf(
			"expected filename config.go, got %s",
			result.Filename,
		)
	}

	if len(result.ChangedLines) != 2 {
		t.Fatalf(
			"expected 2 changed lines, got %d",
			len(result.ChangedLines),
		)
	}

	if result.ChangedLines[0].LineNumber != 12 {
		t.Fatalf(
			"expected first changed line to be 12, got %d",
			result.ChangedLines[0].LineNumber,
		)
	}

	if result.ChangedLines[0].Content !=
		`    password := "hello123"` {

		t.Fatalf(
			"unexpected first changed line: %s",
			result.ChangedLines[0].Content,
		)
	}

	if result.ChangedLines[1].LineNumber != 13 {
		t.Fatalf(
			"expected second changed line to be 13, got %d",
			result.ChangedLines[1].LineNumber,
		)
	}
}