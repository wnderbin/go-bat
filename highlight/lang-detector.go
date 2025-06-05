package highlight

import (
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
)

func DetectFileType(filename, content string) chroma.Lexer {
	lexer := lexers.Match(filename)
	if lexer == nil {
		lexer = lexers.Analyse(content)
	}
	if lexer == nil {
		lexer = lexers.Fallback
	}
	return lexer
}
