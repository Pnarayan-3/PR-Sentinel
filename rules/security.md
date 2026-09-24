# Security Rules

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