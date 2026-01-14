package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/parradam/go-fetch/internal/commands"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 && (args[0] == "help" || args[0] == "--help") {
		printUsage()
		return
	}

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Error: incorrect number of arguments\n\n")
		printUsage()
		os.Exit(1)
	}

	switch args[0] {
	case "fetch":
		parts := strings.Split(args[1], "/")

		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Error: incorrect format for <owner/repo>\n\n")
			printUsage()
			os.Exit(1)
		}

		owner := parts[0]
		repo := parts[1]

		err := commands.Fetch(owner, repo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n\n", args[0])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: go-fetch <command> <args>")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  fetch <owner/repo>  Fetch issues from a GitHub repository")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go-fetch fetch golang/go")
}
