// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JSON

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by JSONParser.
type JSONVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JSONParser#json.
	VisitJson(ctx *JsonContext) interface{}

	// Visit a parse tree produced by JSONParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by JSONParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by JSONParser#members.
	VisitMembers(ctx *MembersContext) interface{}

	// Visit a parse tree produced by JSONParser#member.
	VisitMember(ctx *MemberContext) interface{}

	// Visit a parse tree produced by JSONParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by JSONParser#elements.
	VisitElements(ctx *ElementsContext) interface{}

	// Visit a parse tree produced by JSONParser#str.
	VisitStr(ctx *StrContext) interface{}

	// Visit a parse tree produced by JSONParser#num.
	VisitNum(ctx *NumContext) interface{}
}
