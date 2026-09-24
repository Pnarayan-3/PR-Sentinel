# PR Sentinel Rules

PR Sentinel uses this file to define repository-specific review rules.

---

## Security

### SEC-001 — Hardcoded Password

ID: SEC-001
Category: Security
Severity: CRITICAL
Description: A possible hardcoded password was detected.
Suggestion: Move passwords to environment variables or a secure secret manager.
Pattern: password=

### SEC-002 — Hardcoded API Key

ID: SEC-002
Category: Security
Severity: CRITICAL
Description: A possible hardcoded API key was detected.
Suggestion: Store API keys in environment variables or GitHub Secrets.
Pattern: api_key=

### SEC-003 — Hardcoded Secret

ID: SEC-003
Category: Security
Severity: CRITICAL
Description: A possible hardcoded secret was detected.
Suggestion: Move secrets outside the source code.
Pattern: secret=

---

## Code Quality

### CODE-001 — TODO Marker

ID: CODE-001
Category: Code Quality
Severity: LOW
Description: TODO marker found in source code.
Suggestion: Resolve the TODO or create a tracked issue.
Pattern: TODO

### CODE-002 — Debug Print

ID: CODE-002
Category: Code Quality
Severity: LOW
Description: Debug output was detected.
Suggestion: Remove debug output before merging production code.
Pattern: fmt.Println

---

## Testing

### TEST-001 — Debugger Statement

ID: TEST-001
Category: Testing
Severity: MEDIUM
Description: A debugger statement was detected.
Suggestion: Remove debugger statements before merging.
Pattern: debugger

---

## SQL

### SQL-001 — Possible SQL Injection

ID: SQL-001
Category: Security
Severity: HIGH
Description: SQL construction may contain directly concatenated user input.
Suggestion: Use parameterized queries or prepared statements.
Pattern: SELECT *
