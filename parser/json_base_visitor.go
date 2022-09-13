// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JSON

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseJSONVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJSONVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitMembers(ctx *MembersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitMember(ctx *MemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitElements(ctx *ElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitStr(ctx *StrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitNum(ctx *NumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJSONVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}
