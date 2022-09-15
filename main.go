package main

import (
	"fmt"
	"io"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/hikaru7719/go-json-format/parser"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lexer := parser.NewJSONLexer(antlr.NewInputStream(string(b)))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJSONParser(stream)
	parser.BuildParseTrees = true
	visitor := &TreeVisitor{}
	parser.Json().Accept(visitor)
	fmt.Println(visitor.String())
}
