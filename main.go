package main

import (
	"fmt"
	"os"
	"strings"

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

var _ parser.JSONVisitor = &TreeVisitor{}

var (
	defaultIndent = "    "
)

type TreeVisitor struct {
	parser.BaseJSONVisitor
	builer strings.Builder
	depth  int
}

func (visitor *TreeVisitor) IncrementDepth() {
	visitor.depth += 1
}

func (visitor *TreeVisitor) DecrementDepth() {
	visitor.depth -= 1
}

func (visitor *TreeVisitor) PrintIndent() string {
	return strings.Repeat(defaultIndent, visitor.depth)
}

func (visitor *TreeVisitor) String() string {
	return visitor.builer.String()
}

func (visitor *TreeVisitor) VisitJson(ctx *parser.JsonContext) interface{} {
	ctx.Value().Accept(visitor)
	visitor.builer.WriteString("\n")
	return nil
}

func (visitor *TreeVisitor) VisitValue(ctx *parser.ValueContext) interface{} {
	if t := ctx.Object(); t != nil {
		t.Accept(visitor)
	}

	if t := ctx.Array(); t != nil {
		t.Accept(visitor)
	}

	if t := ctx.Str(); t != nil {
		t.Accept(visitor)
	}

	if t := ctx.Num(); t != nil {
		t.Accept(visitor)
	}

	if t := ctx.Bool_(); t != nil {
		t.Accept(visitor)
	}

	if t := ctx.Nil_(); t != nil {
		t.Accept(visitor)
	}
	return nil
}

func (visitor *TreeVisitor) VisitObject(ctx *parser.ObjectContext) interface{} {
	visitor.builer.WriteString(fmt.Sprintf("%s\n", ctx.LEFT_BRACKET().GetText()))
	visitor.IncrementDepth()

	ctx.Members().Accept(visitor)

	visitor.DecrementDepth()
	visitor.builer.WriteString(fmt.Sprintf("%s%s", visitor.PrintIndent(), ctx.RIGHT_BRACKET().GetText()))
	return nil
}

func (visitor *TreeVisitor) VisitMembers(ctx *parser.MembersContext) interface{} {
	last := len(ctx.AllMember()) - 1
	for i, m := range ctx.AllMember() {
		m.Accept(visitor)
		if i < last {
			visitor.builer.WriteString(ctx.COMMA().GetText())
		}
		visitor.builer.WriteString("\n")
	}
	return nil
}

func (visitor *TreeVisitor) VisitMember(ctx *parser.MemberContext) interface{} {
	visitor.builer.WriteString(fmt.Sprintf("%s", visitor.PrintIndent()))
	ctx.Str().Accept(visitor)
	visitor.builer.WriteString(fmt.Sprintf("%s ", ctx.COLORN().GetText()))
	ctx.Value().Accept(visitor)
	return nil
}

func (visitor *TreeVisitor) VisitArray(ctx *parser.ArrayContext) interface{} {
	visitor.IncrementDepth()
	defer visitor.DecrementDepth()

	ctx.Elements().Accept(visitor)
	return nil
}

func (visitor *TreeVisitor) VisitElements(ctx *parser.ElementsContext) interface{} {
	for _, v := range ctx.AllValue() {
		v.Accept(visitor)
	}
	return nil
}

func (visitor *TreeVisitor) VisitStr(ctx *parser.StrContext) interface{} {
	visitor.builer.WriteString(ctx.GetText())
	return nil
}

func (visitor *TreeVisitor) VisitNum(ctx *parser.NumContext) interface{} {
	visitor.builer.WriteString(ctx.GetText())
	return nil
}

func (visitor *TreeVisitor) VisitBoolean(ctx *parser.BoolContext) interface{} {
	visitor.builer.WriteString(ctx.GetText())
	return nil
}

func (visitor *TreeVisitor) VisitNull(ctx *parser.NilContext) interface{} {
	visitor.builer.WriteString(ctx.GetText())
	return nil
}
