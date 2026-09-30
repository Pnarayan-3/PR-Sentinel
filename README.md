# 🛡️ PR Sentinel

### Automated Pull Request Code Review & Security Analysis Agent

PR Sentinel is a lightweight, rule-based code review agent built in **Go** that integrates with **GitHub Pull Requests through GitHub Actions**.

It analyzes repository code and Pull Request changes against configurable rules, identifies potential issues based on severity, and generates a readable review report.

The project is designed as a foundation for a future automated PR-review agent similar in concept to internal corporate code-review bots.

---

## 🚀 Project Status

The project has reached a functional milestone and is intentionally paused here.

The current implementation can:

* Scan a local repository
* Load configurable review rules
* Analyze source files
* Detect rule-based issues
* Assign severity levels
* Generate Markdown reports
* Run automatically through GitHub Actions
* Detect GitHub Pull Requests
* Retrieve Pull Request changed files 
* Retrieve PR patches/diffs
* Parse PR patches
* Identify added lines from a Pull Request
* Preserve the actual line numbers of added code
* Run the project entirely through GitHub Actions without requiring Go to be installed locally

---

## 🧠 How PR Sentinel Works

At the current stage, PR Sentinel has two main analysis modes.

### 1. Local Repository Analysis
### 2. GitHub Pull Request Analysis

### 🏗️ Current Architecture

```text
                         GitHub
                           │
                           │ Pull Request
                           ▼
                  ┌───────────────────┐
                  │   GitHub Actions  │
                  └─────────┬─────────┘
                            │
                            ▼
                  ┌───────────────────┐
                  │    PR Sentinel    │
                  └─────────┬─────────┘
                            │
                ┌───────────┴───────────┐
                │                       │
                ▼                       ▼
       Local Repository          GitHub PR Mode
                │                       │
                ▼                       ▼
          Repository              GitHub API
            Scanner                    │
                │                       ▼
                │                Changed Files
                │                       │
                │                       ▼
                │                  PR Patch
                │                       │
                │                       ▼
                │                  Diff Parser
                │                       │
                └───────────┬───────────┘
                            │
                            ▼
                     Rule Engine
                            │
                            ▼
                       Findings
                            │
                            ▼
                    Markdown Report
```
---

### ⚙️ Technology Stack

| Technology          | Purpose                   |
| ------------------- | ------------------------- |
| **Go**              | Core agent implementation |
| **GitHub Actions**  | CI/CD execution           |
| **GitHub REST API** | Pull Request integration  |
| **Markdown**        | Rules and reports         |
| **Git**             | Version control           |

---


### 🔄 GitHub Actions Workflow

The Pull Request workflow listens for:

```yaml
on:
  pull_request:
    types:
      - opened
      - synchronize
      - reopened
      - ready_for_review
```

This means the workflow can run when:

* A Pull Request is opened
* New commits are pushed to the Pull Request
* A Pull Request is reopened
* A draft Pull Request becomes ready for review

The workflow currently skips draft Pull Requests:

```yaml
if: github.event.pull_request.draft == false
```

---

### 🚨 Severity Levels

PR Sentinel supports severity-based findings.

Example levels:

```text
🔴 CRITICAL
🟠 HIGH
🟡 MEDIUM
🟢 LOW
```

A finding contains information such as:

```text
Rule ID
Severity
Category
File
Line
Title
Description
Suggestion
```

Example:

```text
🔴 CRITICAL — SEC-001

Hardcoded Password

File:
config.go

Line:
27

Description:
A possible hardcoded password was detected.

Suggestion:
Move passwords to environment variables
or a secure secret manager.
```

---

### 📝 Report Generation

Findings are converted into a Markdown report.

Example:

```markdown
#### 🔴 CRITICAL — SEC-001

**SEC-001 — Hardcoded Password**

📁 `config.go`

📍 Line `27`

**Description:**
A possible hardcoded password was detected.

**Suggestion:**
Move passwords to environment variables
or a secure secret manager.
```

A final review status can also be generated:

```text
❌ FAILED
```

when blocking findings are present.

---

### 💻 Running PR Sentinel

### Local Repository Mode

The application can analyze a repository using the configured repository path and rules.

The general execution is:

```bash
go run ./cmd/sentinel
```

Local Go installation is required for this mode.

---

### 🔬 Patch Parser

The patch parser understands Git diff hunks such as:

```diff
@@ -10,4 +10,6 @@
 func login() {
     username := "admin"
+    password := "hello123"
+    debug := true
 }
```

It identifies added lines while ignoring:

```text
Deleted lines
Diff metadata
No-newline markers
```

The parser also tracks the new file's line numbers.

Example result:

```text
Line 12 → password := "hello123"
Line 13 → debug := true
```

Unit tests were created for this functionality.

---

### 🧪 Testing Strategy

PR Sentinel is designed so that the project can be tested through GitHub Actions.

The CI pipeline can run:

```bash
go test ./...
```

and:

```bash
go build ./...
```

This is particularly useful when developing from a machine where Go is not installed.

The GitHub runner provides the required Go environment.

---

### 🚧 Current Limitations

PR Sentinel is currently a **rule-based prototype/agent**, not a full AI code-review system.

It does not currently provide:

* Natural-language reasoning from an LLM
* Automatic contextual code understanding
* Full semantic analysis
* Automatic inline review comments
* Automatic conversation resolution
* A GitHub App
* Organization-wide installation
* Cross-file reasoning
* Complex vulnerability analysis

These are potential future extensions.

---

### 🔐 Security Considerations

Never commit:

```text
GitHub Personal Access Tokens
API keys
AWS access keys
Cloud credentials
Gemini/OpenAI API keys
Passwords
Private secrets
```

Use:

```text
GitHub Secrets
Environment Variables
GitHub's GITHUB_TOKEN
```

instead.

---

### 📊 Example

A Pull Request introduces:

```diff
 func connect() {
     username := "admin"
+    password := "hello123"
 }
```

PR Sentinel can identify the added line and produce a finding such as:

```text
🔴 CRITICAL — SEC-001

Hardcoded Password

📁 config.go
📍 Line 12

Description:
A possible hardcoded password was detected.

Suggestion:
Move passwords to environment variables
or a secure secret manager.
```

The goal is to catch the issue **during the Pull Request lifecycle**, before the change reaches the main branch.

---

## 👨‍💻 Author

**Pushkar Narayan**
Software Developer | Cloud | Automation 

---
