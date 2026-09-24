# Default PR Sentinel Rules

## Security

### SEC-001 — Hardcoded Password

ID: SEC-001
Category: Security
Severity: CRITICAL
Description: Possible hardcoded password detected.
Suggestion: Use environment variables or a secret manager.
Pattern: password=

### SEC-002 — Hardcoded API Key

ID: SEC-002
Category: Security
Severity: CRITICAL
Description: Possible hardcoded API key detected.
Suggestion: Use environment variables or GitHub Secrets.
Pattern: api_key=

### SEC-003 — Hardcoded Secret

ID: SEC-003
Category: Security
Severity: CRITICAL
Description: Possible hardcoded secret detected.
Suggestion: Store secrets outside source code.
Pattern: secret=

## Quality

### CODE-001 — TODO

ID: CODE-001
Category: Quality
Severity: LOW
Description: TODO marker detected.
Suggestion: Resolve the TODO or create a tracked issue.
Pattern: TODO