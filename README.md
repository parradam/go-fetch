# go-fetch

A CLI tool to fetch GitHub issues and comments.

## Status

🚧 Work in progress

## Current Features

- ✅ Fetch issues from GitHub repositories via REST API
- ✅ Convert GitHub data to internal domain models

## Planned Features

- CLI interface with flags for repo selection
- Export to Markdown/HTML
- Support multiple repositories
- Track sync state to avoid re-fetching
- Fetch comments for issues

## Development

Run:

```bash
go run cmd/go-fetch/main.go
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