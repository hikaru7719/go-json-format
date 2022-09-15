package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/hikaru7719/go-json-format/parser"
)

func main() {
	b, err := os.ReadFile("test.json")
	if err != nil {
		panic(err)
	}
	lexer := parser.NewJSONLexer(antlr.NewInputStream(string(b)))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJSONParser(stream)
	parser.BuildParseTrees = true
	visitor := &TreeVisitor{}
	parser.Json().Accept(visitor)
	fmt.Println(visitor.String())
}
