# 🔐 Security Policy

Security is a priority. We maintain a proactive stance to identify and fix vulnerabilities in **go-sanitize**.

<br/>

## 🛠️ Supported & Maintained Versions

| Version | Status               |
|---------|----------------------|
| 1.x.x   | ✅ Supported & Active |

<br/>

## 📨 Reporting a Vulnerability

If you’ve found a security issue, **please don’t open a public issue or PR**.

Instead, send a private email to:
📧 [go-sanitize@mrz1818.com](mailto:go-sanitize@mrz1818.com)

Include the following:

* 🕵️ Description of the issue and its impact
* 🧪 Steps to reproduce or a working PoC
* 🔧 Any known workarounds or mitigations

We welcome responsible disclosures from researchers, vendors, users, and curious tinkerers alike.

<br/>

## 📅 What to Expect

* 🧾 **Acknowledgment** within 72 hours
* 📢 **Status updates** every 5 business days
* ✅ **Resolution target** of 30 days (for confirmed vulnerabilities)

Prefer encrypted comms? Let us know in your initial email—we’ll reply with our PGP public key. 
All official security responses are signed with it.

<br/>

## 🧪 Security Tooling

We regularly scan for known vulnerabilities using:

* [`govulncheck`](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)

Want to run it yourself?

```sh
make govulncheck
```

This will check your local build for known issues in Go modules.

<br/>

## 🛡️ Security Standards

We follow the [OpenSSF](https://openssf.org) best practices to ensure this repository remains compliant with industry‑standard open source security guidelines.

<br/>

## 🛠️ GitHub Security Workflows

To proactively protect this repository, we use several automated GitHub workflows:

- **[CodeQL Analysis](./workflows/codeql-analysis.yml)**: Scans the codebase for security vulnerabilities and coding errors using GitHub's CodeQL engine on every push and pull request to the `master` branch.
- **[Gitleaks Scan](./workflows/check-for-leaks.yml)**: Runs daily and on demand to detect secrets or sensitive data accidentally committed to the repository, helping prevent credential leaks.
- **[OpenSSF Scorecard](./workflows/scorecard.yml)**: Periodically evaluates the repository against OpenSSF Scorecard checks, providing insights and recommendations for improving supply chain security and best practices.

These workflows help us identify, remediate, and prevent security issues as early as possible in the development lifecycle. For more details, see the workflow files in the [`.github/workflows/`](https://github.com/mrz1818/go-sanitize/tree/master/.github/workflows) directory.

<br/>
