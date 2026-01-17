# go-fetch

A CLI tool to fetch GitHub issues and save them as Markdown files.

## Status

🚧 Work in progress

## Current Features

- ✅ Fetch issues from GitHub via REST API (currently first 30 issues)
- ✅ Export to Markdown format
- ✅ CLI interface with argument parsing and validation
- ✅ User-friendly error handling
- ✅ Convert GitHub data to internal domain models

## Planned Features

- Export to HTML format
- Support multiple repositories via config file
- Track sync state to avoid re-fetching
- Fetch comments for issues
- Pagination support for larger repositories (currently fetches first 30 issues)
- Custom formatting

## Usage

Fetch issues from a GitHub repository and save to Markdown:

```bash
go-fetch fetch golang/go
# Creates golang-go-issues.md
```

Show help:

```bash
go-fetch --help
```

## Output

Issues are saved to `{owner}-{repo}-issues.md` in the current directory with:
- Repository title and fetch timestamp
- Issue number, title, author, state, and creation date
- Full issue body text
- Markdown formatting for easy reading

Example output structure:
```markdown
# Issues for golang/go

Fetched: 2025-01-17T20:30:00Z

## Issue #12345: Bug in runtime

**Author:** username  
**State:** open  
**Created:** 2025-01-15

Issue description here...

---
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

## Project Structure

```
go-fetch
├── LICENSE
├── README.md
├── cmd
│   └── go-fetch
│       └── main.go            # entry point
├── go-fetch
├── go.mod
└── internal
    ├── api                    # external API clients
    │   ├── client.go
    │   └── github
    │       ├── client.go
    │       └── types.go
    ├── commands
    │   └── fetch.go
    ├── models                 # domain models
    │   ├── comment.go
    │   └── issue.go
    └── output
        ├── formatter.go
        └── markdown.go
```

## License

See the [LICENSE](LICENSE) file.