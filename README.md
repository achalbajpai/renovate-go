## Task Submission
This project is a minimal Go web server built with the Gin framework, intentionally using a vulnerable version to demonstrate Renovate's automated security dependency updates.

-----

### Overview

The server provides basic HTTP endpoints and highlights how Renovate can detect and propose fixes for known security vulnerabilities in project dependencies.

-----

### Features

  * Go Web Server
  * **Demonstrates automated vulnerability remediation with Renovate**
<img width="1052" alt="Screenshot 2025-05-28 at 12 27 43 PM" src="https://github.com/user-attachments/assets/58f30314-208d-44e6-87ea-207fe4e8824a" />

-----

### Vulnerable Dependency

**`github.com/gin-gonic/gin v1.6.3`**

This version is affected by **CVE-2020-28483**, a denial-of-service vulnerability related to deeply nested JSON requests.

<img width="1759" alt="Screenshot 2025-05-28 at 12 01 11 PM" src="https://github.com/user-attachments/assets/7c66f91f-4dbe-4844-acdd-9d995dd7fde4" />

-----

## Quick Start

### Prerequisites

  * Go 1.23.4 or later
  * Git

### Running the Server

```bash
go run main.go
```

The server will be accessible on `http://localhost:8080`.


### Test Usage

```Endpoints
 http://localhost:8080/
 http://localhost:8080/health
 http://localhost:8080/file/test.txt
```

-----

## Renovate Configuration

 `renovate.json` file configured to:

| Feature                                 | Description |
|-----------------------------------------|-------------|
| **Vulnerability-Only Updates**          | Disables all regular Go dependency updates (major, minor, patch) while specifically enabling vulnerability alerts. This ensures dependencies are updated **only** when security issues are detected. |
| **Immediate Security Response**         | Configured for instant pull request creation when vulnerabilities are found — with no minimum release age, scheduling delays, or dependency dashboard approval required. This ensures fast resolution of critical security issues. |
| **Security-Focused Labeling & Branching** | Automatically applies `"security"` and `"vulnerability"` labels to pull requests, appends `[SECURITY]` to commit messages, and creates dedicated, vulnerability-specific branch names for easy tracking and clarity. |
| **Go Module Optimization**              | Tailored for Go projects using the `"lowest"` vulnerability fix strategy, ensuring minimum breaking changes while applying essential security patches. Configured for seamless operation across all time zones. |



-----

## Project Structure

```
.
├── main.go            # Main application server
├── go.mod             # Go module definition
├── go.sum             # Dependency checksums
├── renovate.json      # Renovate configuration
└── README.md          # Project README
```

-----

## Purpose

Antrea Renovate Proposal #7155

This project was created for #7155, which demonstrates usage of Renovate for a Go project.

_The implementation follows Go module layout best practices and intentionally uses Gin v1.6.3 to trigger Renovate's vulnerability detection system._
