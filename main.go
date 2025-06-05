package main

import (
	"fmt"
	"go-bat/highlight"
	"os"

	"github.com/alecthomas/chroma/quick"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./go-bat <file>")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
	defer file.Close()
	content := fmt.Sprintf("%s\n", readFile(file))
	lexer := highlight.DetectFileType(file.Name(), content)
	err = quick.Highlight(os.Stdout, content, lexer.Config().Name, "terminal16m", "monokai")
	if err != nil {
		fmt.Printf("Highlight error: %v\n", err)
	}
}

func readFile(file *os.File) string {
	content, _ := os.ReadFile(file.Name())
	return string(content)
}
