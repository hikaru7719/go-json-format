package main

import (
	"fmt"
	"strings"

	"github.com/hikaru7719/go-json-format/parser"
)

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
	visitor.builer.WriteString(fmt.Sprintf("%s", visitor.PrintIndent()))
	ctx.Member().Accept(visitor)

	if ctx.COMMA() == nil {
		visitor.builer.WriteString("\n")
		return nil
	}
	visitor.builer.WriteString(ctx.COMMA().GetText())
	visitor.builer.WriteString("\n")

	ctx.Members().Accept(visitor)
	return nil
}

func (visitor *TreeVisitor) VisitMember(ctx *parser.MemberContext) interface{} {
	ctx.Str().Accept(visitor)
	visitor.builer.WriteString(fmt.Sprintf("%s ", ctx.COLORN().GetText()))
	ctx.Value().Accept(visitor)
	return nil
}

func (visitor *TreeVisitor) VisitArray(ctx *parser.ArrayContext) interface{} {
	visitor.builer.WriteString(fmt.Sprintf("%s\n", ctx.LEFT_SQUARE_BRACKET().GetText()))
	visitor.IncrementDepth()

	ctx.Elements().Accept(visitor)

	visitor.DecrementDepth()
	visitor.builer.WriteString(fmt.Sprintf("%s%s", visitor.PrintIndent(), ctx.RIGHT_SQUARE_BRACKET().GetText()))
	return nil
}

func (visitor *TreeVisitor) VisitElements(ctx *parser.ElementsContext) interface{} {
	visitor.builer.WriteString(fmt.Sprintf("%s", visitor.PrintIndent()))
	ctx.Value().Accept(visitor)

	if ctx.COMMA() == nil {
		visitor.builer.WriteString("\n")
		return nil
	}
	visitor.builer.WriteString(ctx.COMMA().GetText())
	visitor.builer.WriteString("\n")

	ctx.Elements().Accept(visitor)
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

func (visitor *TreeVisitor) VisitBool(ctx *parser.BoolContext) interface{} {
	visitor.builer.WriteString(ctx.GetText())
	return nil
}

func (visitor *TreeVisitor) VisitNil(ctx *parser.NilContext) interface{} {
	visitor.builer.WriteString(ctx.GetText())
	return nil
}
