# 🤖 PR Sentinel

PR Sentinel is a rule-based code review agent written in Go.

It can analyze source code and identify configurable security,
quality, testing, and performance issues.

## Features

- Rule-based code analysis
- Pull request analysis
- GitHub Actions integration
- Markdown reports
- Configurable repository rules
- Severity-based findings
- Local repository analysis
- Docker support
- No external AI API required

## Architecture

```text
Pull Request
     |
     v
GitHub Actions
     |
     v
Go PR Sentinel
     |
     v
Rule Engine
     |
     v
PR-SENTINEL.md
     |
     v
Findings
     |
     v
GitHub PR Comment