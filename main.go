package main

import (
	"flag"
	"fmt"
	"go-bat/highlight"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
)

func main() {
	showGit := flag.Bool("git", false, "Show git changes")
	lineNumbers := flag.Bool("n", false, "Show line numbers")
	version := flag.Bool("version", false, "Show version")
	help := flag.Bool("help", false, "Instructions for use")
	flag.Parse()

	if *version {
		fmt.Println("\x1b[36m GoBat v0.9.1 [BETA] \x1b[0m")
		return
	}
	if *help {
		fmt.Println("\x1b[1mHelp:\x1b[0m\n \x1b[1m* Launch flags:\x1b[0m\n\t\x1b[1m--git:\x1b[0m Highlighting git changes in output.\n\t\x1b[1m-n:\x1b[0m Line numbering\n\t\x1b[1m--version:\x1b[0m GoBat version\n\t\x1b[1m--help:\x1b[0m Help-guide")
		fmt.Println("\x1b[1m * Use cases:\x1b[0m\n\t1. View file contents with syntax highlighting:\n\t ./go-bat <filepath>\n\t2. View git changes\n\t ./go-bat --git <filepath>\n\t3. Line numbering\n\t ./go-bat -n <filepath>\n\t4. GoBat version:\n\t ./go-bat --version")
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: no file specified")
		fmt.Println("Usage: ./go-bat [flags] <file>")
		os.Exit(1)
	}

	filePath := args[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	content := highlight.ReadFile(file)

	if *showGit {
		changedLines, err := highlight.GetGitDiff(filePath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		highlight.PrintWithGitHighlighting(content, changedLines)
		return
	}

	if *lineNumbers {
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			fmt.Printf("\x1b[33m%6d\x1b[0m %s\n", i+1, line)
		}
		return
	}

	fullContent := fmt.Sprintf("%s\n", content)
	lexer := highlight.DetectFileType(filePath, fullContent)
	err = quick.Highlight(os.Stdout, fullContent, lexer.Config().Name, "terminal16m", "monokai")
	if err != nil {
		fmt.Printf("Highlight error: %v\n", err)
		os.Exit(1)
	}
}
