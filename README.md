# go-fetch

A CLI tool to fetch GitHub issues and comments.

## Status

🚧 Work in progress

## Current Features

- ✅ Fetch issues from GitHub repositories via REST API
- ✅ CLI interface with argument parsing and validation
- ✅ Built-in error handling
- ✅ Convert GitHub data to internal domain models

## Planned Features

- Export to Markdown/HTML
- Support multiple repositories via config file
- Track sync state to avoid re-fetching
- Fetch comments for issues
- Pagination support for larger repositories

## Usage

Fetch issues from a GitHub repository:

```bash
go-fetch fetch golang/go
```

Show help:

```bash
go-fetch --help
```

## Development

Run:

```bash
go run cmd/go-fetch/main.go fetch golang/go
```

Format code:

```bash
go fmt ./...
```

Lint code:

```bash
golangci-lint run
```

Build binary:

```bash
go build -o go-fetch cmd/go-fetch/main.go
```

Run binary:

```bash
./go-fetch fetch golang/go
```

## Project structure

```
go-fetch
├── cmd
│   └── go-fetch
│       └── main.go             # entry point
├── go.mod
├── internal
│   ├── api                     # external API clients
│   │   ├── client.go
│   │   └── github
│   │       ├── client.go
│   │       └── types.go
│   └── models                  # domain models
│       ├── comment.go
│       └── issue.go
└── README.md
```