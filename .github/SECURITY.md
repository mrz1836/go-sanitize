# Security Policy

## Supported & Maintained Versions

| Version | Supported          |
|---------|--------------------|
| 1.x.x   | :white_check_mark: |

## Reporting a Vulnerability

Please email the [project maintainers](mailto:go-sanitize@mrz1818.com) with a detailed
report. We welcome submissions from independent researchers, industry
organizations, vendors and customers. Include:

- A description of the issue and its impact
- Steps to reproduce or proof of concept
- Known workarounds or mitigation

Do **not** open a public issue or pull request for security matters.

### Response Expectations

- **Acknowledgment** within 72 hours of receipt
- **Status updates** at least every 5 business days
- **Resolution target** of 30 days for confirmed vulnerabilities

If you prefer encrypted communication, request our PGP public key in your
initial email. Official security correspondence will be signed with this key.

### Security Tooling

We use [`govulncheck`](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck) to
scan dependencies. You can run `make govulncheck` locally to check your own
builds.

